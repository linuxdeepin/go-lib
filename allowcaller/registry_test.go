// SPDX-FileCopyrightText: 2026 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package allowcaller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/godbus/dbus/v5"
)

// Test scope strings stand in for the v25 D-Bus object identifiers that the
// consuming projects define. The package itself is scope-name agnostic.
const (
	testDaemonScope = "org.deepin.dde.Daemon1:/org/deepin/dde/Daemon1"
	testPowerScope  = "org.deepin.dde.Power1:/org/deepin/dde/Power1"

	testPrivilegedGroupID = 996
)

type fakeBusService struct {
	owners  map[string]bool
	uids    map[string]uint32
	groups  map[string][]uint32
	pids    map[string]uint32
	parents map[uint32]uint32
	busID   string
}

func (s *fakeBusService) NameHasOwner(name string) (bool, error) {
	return s.owners[name], nil
}

func (s *fakeBusService) GetConnUID(name string) (uint32, error) {
	return s.uids[name], nil
}

func (s *fakeBusService) GetConnPID(name string) (uint32, error) {
	pid, ok := s.pids[name]
	if !ok {
		return 0, fmt.Errorf("PID for %s is unavailable", name)
	}
	return pid, nil
}

func (s *fakeBusService) GetConnGroups(name string) ([]uint32, error) {
	return s.groups[name], nil
}

func (s *fakeBusService) GetBusID() (string, error) {
	return s.busID, nil
}

func newTestRegistry(service *fakeBusService, stateFile string, privilegedGroupID uint32) *Registry {
	registry := newRegistry(service, stateFile, privilegedGroupID)
	registry.processParent = func(pid uint32) (uint32, error) {
		parentPID, ok := service.parents[pid]
		if !ok {
			return 0, fmt.Errorf("parent PID for %d is unavailable", pid)
		}
		return parentPID, nil
	}
	return registry
}

func TestRegistryAuthorize(t *testing.T) {
	stateFile := filepath.Join(t.TempDir(), "allow-callers.json")
	service := &fakeBusService{
		owners:  map[string]bool{":1.1": true, ":1.10": true, ":1.11": true, ":1.12": true},
		uids:    map[string]uint32{":1.1": 1000, ":1.10": 1000, ":1.11": 1000, ":1.12": 0},
		groups:  map[string][]uint32{":1.1": {testPrivilegedGroupID}},
		pids:    map[string]uint32{":1.1": 101, ":1.10": 110},
		parents: map[uint32]uint32{110: 101},
		busID:   "bus-a",
	}
	registry := newTestRegistry(service, stateFile, testPrivilegedGroupID)

	if err := registry.AddCaller(testDaemonScope, dbus.Sender(":1.1"), ":1.10"); err != nil {
		t.Fatal(err)
	}
	if err := registry.Authorize(testDaemonScope, dbus.Sender(":1.10")); err != nil {
		t.Fatalf("registered caller was denied: %v", err)
	}
	if err := registry.Authorize(testDaemonScope, dbus.Sender(":1.11")); err == nil {
		t.Fatal("same-UID unregistered caller was allowed")
	}
	if err := registry.Authorize(testPowerScope, dbus.Sender(":1.10")); err == nil {
		t.Fatal("caller registered for another scope was allowed")
	}
	if err := registry.Authorize(testPowerScope, dbus.Sender(":1.12")); err != nil {
		t.Fatalf("root caller was denied: %v", err)
	}
}

func TestRegistryPersistenceAndRemoval(t *testing.T) {
	stateFile := filepath.Join(t.TempDir(), "allow-callers.json")
	service := &fakeBusService{
		owners: map[string]bool{":1.20": true, ":1.21": false},
		uids:   map[string]uint32{":1.20": 1000},
		groups: make(map[string][]uint32),
		busID:  "bus-a",
	}
	content, err := json.Marshal(persistedState{
		BusID: "bus-a",
		Callers: map[string][]string{
			testDaemonScope: {":1.20", ":1.21"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(stateFile, content, 0600); err != nil {
		t.Fatal(err)
	}

	registry := newTestRegistry(service, stateFile, testPrivilegedGroupID)
	if err := registry.load(); err != nil {
		t.Fatal(err)
	}
	if err := registry.Authorize(testDaemonScope, dbus.Sender(":1.20")); err != nil {
		t.Fatalf("live persisted caller was denied: %v", err)
	}
	registry.RemoveCaller(":1.20")
	if err := registry.Authorize(testDaemonScope, dbus.Sender(":1.20")); err == nil {
		t.Fatal("removed caller was allowed")
	}
}

func TestRegistryRejectsStateFromAnotherBus(t *testing.T) {
	stateFile := filepath.Join(t.TempDir(), "allow-callers.json")
	content, err := json.Marshal(persistedState{
		BusID: "old-bus",
		Callers: map[string][]string{
			testDaemonScope: {":1.30"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(stateFile, content, 0600); err != nil {
		t.Fatal(err)
	}

	service := &fakeBusService{
		owners: map[string]bool{":1.30": true},
		uids:   map[string]uint32{":1.30": 1000},
		groups: make(map[string][]uint32),
		busID:  "new-bus",
	}
	registry := newTestRegistry(service, stateFile, testPrivilegedGroupID)
	if err := registry.load(); err != nil {
		t.Fatal(err)
	}
	if err := registry.Authorize(testDaemonScope, dbus.Sender(":1.30")); err == nil {
		t.Fatal("caller restored from a different system-bus lifetime")
	}
}

func TestRegistryConcurrentAddsPersistAllCallers(t *testing.T) {
	stateFile := filepath.Join(t.TempDir(), "allow-callers.json")
	service := &fakeBusService{
		owners:  map[string]bool{":1.1": true},
		uids:    map[string]uint32{":1.1": 1000},
		groups:  map[string][]uint32{":1.1": {testPrivilegedGroupID}},
		pids:    map[string]uint32{":1.1": 101},
		parents: make(map[uint32]uint32),
		busID:   "bus-a",
	}
	const callerCount = 20
	for i := 0; i < callerCount; i++ {
		name := ":1." + strconv.Itoa(100+i)
		service.owners[name] = true
		service.uids[name] = 1000
		service.pids[name] = uint32(200 + i)
		service.parents[uint32(200+i)] = 101
	}

	registry := newTestRegistry(service, stateFile, testPrivilegedGroupID)
	errCh := make(chan error, callerCount)
	var wg sync.WaitGroup
	for i := 0; i < callerCount; i++ {
		name := ":1." + strconv.Itoa(100+i)
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			errCh <- registry.AddCaller(testDaemonScope, dbus.Sender(":1.1"), name)
		}(name)
	}
	wg.Wait()
	close(errCh)
	for err := range errCh {
		if err != nil {
			t.Fatal(err)
		}
	}

	content, err := ioutil.ReadFile(stateFile)
	if err != nil {
		t.Fatal(err)
	}
	var state persistedState
	if err := json.Unmarshal(content, &state); err != nil {
		t.Fatal(err)
	}
	if got := len(state.Callers[testDaemonScope]); got != callerCount {
		t.Fatalf("persisted caller count = %d, want %d", got, callerCount)
	}
}

func TestRegistryValidatesRegistrar(t *testing.T) {
	tests := []struct {
		name       string
		sender     dbus.Sender
		target     string
		uids       map[string]uint32
		groups     map[string][]uint32
		descendant bool
		wantError  bool
	}{
		{
			name:       "privileged same uid descendant",
			sender:     dbus.Sender(":1.1"),
			target:     ":1.10",
			uids:       map[string]uint32{":1.1": 1000, ":1.10": 1000},
			groups:     map[string][]uint32{":1.1": {testPrivilegedGroupID}},
			descendant: true,
		},
		{
			name:      "same uid unrelated process",
			sender:    dbus.Sender(":1.5"),
			target:    ":1.50",
			uids:      map[string]uint32{":1.5": 1000, ":1.50": 1000},
			groups:    map[string][]uint32{":1.5": {testPrivilegedGroupID}},
			wantError: true,
		},
		{
			name:      "missing privileged group",
			sender:    dbus.Sender(":1.2"),
			target:    ":1.20",
			uids:      map[string]uint32{":1.2": 1000, ":1.20": 1000},
			groups:    map[string][]uint32{":1.2": {1000}},
			wantError: true,
		},
		{
			name:      "cross uid target",
			sender:    dbus.Sender(":1.3"),
			target:    ":1.30",
			uids:      map[string]uint32{":1.3": 1000, ":1.30": 1001},
			groups:    map[string][]uint32{":1.3": {testPrivilegedGroupID}},
			wantError: true,
		},
		{
			name:   "root cross uid target",
			sender: dbus.Sender(":1.4"),
			target: ":1.40",
			uids:   map[string]uint32{":1.4": 0, ":1.40": 1001},
			groups: make(map[string][]uint32),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &fakeBusService{
				owners:  map[string]bool{tt.target: true},
				uids:    tt.uids,
				groups:  tt.groups,
				pids:    map[string]uint32{string(tt.sender): 101, tt.target: 110},
				parents: map[uint32]uint32{110: 1, 1: 0},
				busID:   "bus-a",
			}
			if tt.descendant {
				service.parents[110] = 101
			}
			registry := newTestRegistry(service, filepath.Join(t.TempDir(), "allow-callers.json"), testPrivilegedGroupID)
			err := registry.AddCaller(testDaemonScope, tt.sender, tt.target)
			if tt.wantError && err == nil {
				t.Fatal("expected registration to be denied")
			}
			if !tt.wantError && err != nil {
				t.Fatalf("registration was denied: %v", err)
			}
		})
	}
}

func TestAddCallerNotVisibleBeforePersistenceSucceeds(t *testing.T) {
	service := &fakeBusService{
		owners:  map[string]bool{":1.1": true, ":1.50": true},
		uids:    map[string]uint32{":1.1": 1000, ":1.50": 1000},
		groups:  map[string][]uint32{":1.1": {testPrivilegedGroupID}},
		pids:    map[string]uint32{":1.1": 101, ":1.50": 150},
		parents: map[uint32]uint32{150: 101},
		busID:   "bus-a",
	}
	registry := newTestRegistry(service, filepath.Join(t.TempDir(), "allow-callers.json"), testPrivilegedGroupID)
	persistStarted := make(chan struct{})
	releasePersist := make(chan struct{})
	registry.persistState = func(persistedState) error {
		close(persistStarted)
		<-releasePersist
		return errors.New("persist failed")
	}

	addResult := make(chan error, 1)
	go func() {
		addResult <- registry.AddCaller(testDaemonScope, dbus.Sender(":1.1"), ":1.50")
	}()
	<-persistStarted

	authorizeResult := make(chan error, 1)
	go func() {
		authorizeResult <- registry.Authorize(testDaemonScope, dbus.Sender(":1.50"))
	}()
	select {
	case err := <-authorizeResult:
		t.Fatalf("authorization completed before registration commit: %v", err)
	case <-time.After(50 * time.Millisecond):
	}

	close(releasePersist)
	if err := <-addResult; err == nil {
		t.Fatal("expected persistence failure")
	}
	if err := <-authorizeResult; err == nil {
		t.Fatal("failed registration became authorized")
	}
}

func TestIsProcessDescendant(t *testing.T) {
	parents := map[uint32]uint32{400: 300, 300: 200, 200: 1, 1: 0}
	processParent := func(pid uint32) (uint32, error) {
		parentPID, ok := parents[pid]
		if !ok {
			return 0, fmt.Errorf("parent PID for %d is unavailable", pid)
		}
		return parentPID, nil
	}

	isDescendant, err := isProcessDescendant(400, 200, processParent)
	if err != nil {
		t.Fatal(err)
	}
	if !isDescendant {
		t.Fatal("multi-level descendant was rejected")
	}

	isDescendant, err = isProcessDescendant(400, 999, processParent)
	if err != nil {
		t.Fatal(err)
	}
	if isDescendant {
		t.Fatal("unrelated process was accepted as a descendant")
	}

	cyclicParent := func(pid uint32) (uint32, error) {
		if pid == 10 {
			return 11, nil
		}
		return 10, nil
	}
	if _, err := isProcessDescendant(10, 99, cyclicParent); err == nil {
		t.Fatal("cyclic process ancestry was accepted")
	}
}

func TestGetProcessParentPID(t *testing.T) {
	parentPID, err := getProcessParentPID(uint32(os.Getpid()))
	if err != nil {
		t.Fatal(err)
	}
	if want := uint32(os.Getppid()); parentPID != want {
		t.Fatalf("parent PID = %d, want %d", parentPID, want)
	}
}

func TestGetProcessGroupsIncludesEffectiveGroup(t *testing.T) {
	groups, err := getProcessGroups(uint32(os.Getpid()))
	if err != nil {
		t.Fatal(err)
	}
	if !containsGroup(groups, uint32(os.Getegid())) {
		t.Fatalf("effective gid %d was not found in process credentials %v", os.Getegid(), groups)
	}
}
