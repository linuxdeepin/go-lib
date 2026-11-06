// SPDX-FileCopyrightText: 2026 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

// Package allowcaller implements the deepin-security-loader caller allowlist
// mechanism shared by the D-Bus services that are launched by
// deepin-security-loader.
//
// A service that owns protected D-Bus methods keeps a Registry of the exact
// system-bus unique names that deepin-security-loader has vouched for. Each
// protected method calls Authorize(scope, sender) and rejects any caller whose
// unique name was not registered for the scope via AddCaller.
//
// The registry is multi-scope: a single Registry can protect several D-Bus
// objects (each identified by an arbitrary scope string), so one process can
// serve several protected interfaces with a shared persisted state file.
//
// This is a port of the v20 dde-daemon securityloader/allowcaller.go and
// dde-api locale-helper/allow_caller.go implementations, unified and made
// reusable. The scope constants (v25 D-Bus names) live in the consuming
// projects; this package is scope-name agnostic.
package allowcaller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/log"
)

const (
	// privilegedGroup is the unix group whose members (together with root)
	// are trusted to register a caller on behalf of a process they spawned.
	privilegedGroup = "deepin-daemon"
	// invalidGroupID sentinel used when the privileged group cannot be
	// resolved; registrations from non-root senders are then rejected.
	invalidGroupID = ^uint32(0)

	// nameOwnerChangedMatchRule is the D-Bus match rule used to learn when a
	// registered unique name disappears from the bus so its entry can be
	// dropped without waiting for the next process restart.
	nameOwnerChangedMatchRule = "type='signal'," +
		"interface='org.freedesktop.DBus'," +
		"member='NameOwnerChanged'," +
		"path='/org/freedesktop/DBus'"
)

var logger = log.NewLogger("go-lib/allowcaller")

// busService is the subset of *dbusutil.Service the Registry depends on, so
// the core logic can be exercised with a fake bus in tests.
type busService interface {
	NameHasOwner(name string) (bool, error)
	GetConnUID(name string) (uint32, error)
	GetConnPID(name string) (uint32, error)
	GetConnGroups(name string) ([]uint32, error)
	GetBusID() (string, error)
}

// serviceBus adapts a *dbusutil.Service to busService. GetConnGroups is
// derived from the connection PID plus the process credentials in /proc, so
// consumers only need the standard dbusutil.Service methods.
type serviceBus struct {
	*dbusutil.Service
}

func (s serviceBus) GetBusID() (string, error) {
	var id string
	err := s.Conn().BusObject().Call("org.freedesktop.DBus.GetId", 0).Store(&id)
	return id, err
}

func (s serviceBus) GetConnGroups(name string) ([]uint32, error) {
	pid, err := s.GetConnPID(name)
	if err != nil {
		return nil, err
	}
	return getProcessGroups(pid)
}

// persistedState is the on-disk representation of the registry. The BusID
// scopes the state to the current system-bus lifetime: a fresh bus daemon
// instance gets a new BusID, so stale entries from a previous boot are
// ignored instead of being trusted.
type persistedState struct {
	BusID   string              `json:"busId"`
	Callers map[string][]string `json:"callers"`
}

// Registry stores the exact system-bus unique names authorized by
// deepin-security-loader for each exported D-Bus object (scope).
type Registry struct {
	service           busService
	stateFile         string
	busID             string
	privilegedGroupID uint32
	processParent     func(uint32) (uint32, error)

	mu        sync.RWMutex
	callers   map[string]map[string]struct{}
	persistMu sync.Mutex

	// persistState, when set, replaces writeState. Tests use it to observe
	// or block the persistence step without touching the filesystem.
	persistState func(persistedState) error

	signalLoop *dbusutil.SignalLoop
	matchRule  string
}

// NewRegistry creates a Registry backed by service and persists its state to
// stateFile. It wires a NameOwnerChanged listener so entries for vanished bus
// names are removed automatically. Failures are logged as warnings and the
// returned registry is still usable: a missing or corrupt state file simply
// means the registry starts empty, matching the daemon's must-start behavior.
func NewRegistry(service *dbusutil.Service, stateFile string) *Registry {
	groupID, err := lookupGroupID(privilegedGroup)
	if err != nil {
		logger.Warningf("failed to resolve privileged group %s: %v", privilegedGroup, err)
		groupID = invalidGroupID
	}
	registry := newRegistry(serviceBus{service}, stateFile, groupID)
	if err := registry.load(); err != nil {
		logger.Warning("failed to load allow-caller state:", err)
	}

	registry.signalLoop = dbusutil.NewSignalLoop(service.Conn(), 10)
	registry.signalLoop.Start()
	if err := registry.watchOwnerChanges(); err != nil {
		logger.Warning("failed to watch allow-caller owner changes:", err)
	}

	return registry
}

// newRegistry builds the testable core: no signal loop, no owner-change watch.
// processParent defaults to reading /proc and can be overridden by tests.
func newRegistry(service busService, stateFile string, privilegedGroupID uint32) *Registry {
	busID, err := service.GetBusID()
	if err != nil {
		logger.Warning("failed to get system bus ID:", err)
	}
	return &Registry{
		service:           service,
		stateFile:         stateFile,
		busID:             busID,
		privilegedGroupID: privilegedGroupID,
		processParent:     getProcessParentPID,
		callers:           make(map[string]map[string]struct{}),
	}
}

func lookupGroupID(name string) (uint32, error) {
	group, err := user.LookupGroup(name)
	if err != nil {
		return 0, err
	}
	value, err := strconv.ParseUint(group.Gid, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid gid %q for group %s: %w", group.Gid, name, err)
	}
	return uint32(value), nil
}

func getProcessGroups(pid uint32) ([]uint32, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/status", pid))
	if err != nil {
		return nil, err
	}

	var groups []uint32
	for _, line := range strings.Split(string(content), "\n") {
		if !strings.HasPrefix(line, "Gid:") && !strings.HasPrefix(line, "Groups:") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		for _, value := range strings.Fields(parts[1]) {
			gid, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				return nil, fmt.Errorf("invalid process gid %q: %w", value, err)
			}
			groups = append(groups, uint32(gid))
		}
	}
	if len(groups) == 0 {
		return nil, fmt.Errorf("no group credentials found for pid %d", pid)
	}
	return groups, nil
}

func getProcessParentPID(pid uint32) (uint32, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/status", pid))
	if err != nil {
		return 0, err
	}

	for _, line := range strings.Split(string(content), "\n") {
		if !strings.HasPrefix(line, "PPid:") {
			continue
		}
		fields := strings.Fields(strings.TrimPrefix(line, "PPid:"))
		if len(fields) != 1 {
			return 0, fmt.Errorf("invalid PPid entry for pid %d", pid)
		}
		parentPID, err := strconv.ParseUint(fields[0], 10, 32)
		if err != nil {
			return 0, fmt.Errorf("invalid parent pid %q for pid %d: %w", fields[0], pid, err)
		}
		return uint32(parentPID), nil
	}
	return 0, fmt.Errorf("no PPid entry found for pid %d", pid)
}

// IsProcessDescendant reports whether pid is a descendant of ancestorPID.
func IsProcessDescendant(pid, ancestorPID uint32) (bool, error) {
	return isProcessDescendant(pid, ancestorPID, getProcessParentPID)
}

func isProcessDescendant(pid, ancestorPID uint32, processParent func(uint32) (uint32, error)) (bool, error) {
	if pid == 0 || ancestorPID == 0 || pid == ancestorPID {
		return false, nil
	}

	visited := make(map[uint32]struct{})
	currentPID := pid
	for currentPID != 0 {
		if _, exists := visited[currentPID]; exists {
			return false, fmt.Errorf("cycle detected in process ancestry at pid %d", currentPID)
		}
		visited[currentPID] = struct{}{}

		parentPID, err := processParent(currentPID)
		if err != nil {
			return false, err
		}
		if parentPID == ancestorPID {
			return true, nil
		}
		currentPID = parentPID
	}
	return false, nil
}

// Close releases the owner-change watch and signal loop. It is safe to call
// on a nil registry or one created with newRegistry (no signal loop).
func (r *Registry) Close() {
	if r == nil {
		return
	}
	if r.signalLoop != nil {
		if r.matchRule != "" && r.service != nil {
			// best-effort; the connection owns the subscription for its
			// lifetime anyway.
			if bus, ok := r.service.(serviceBus); ok {
				_ = bus.Conn().BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, r.matchRule).Err
			}
		}
		r.signalLoop.Stop()
		r.signalLoop = nil
	}
}

// watchOwnerChanges subscribes to NameOwnerChanged and drops registry entries
// for unique names that vanish from the bus. It is set up by NewRegistry.
func (r *Registry) watchOwnerChanges() error {
	bus, ok := r.service.(serviceBus)
	if !ok {
		return errors.New("owner-change watch requires a *dbusutil.Service-backed registry")
	}
	conn := bus.Conn()
	if err := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, nameOwnerChangedMatchRule).Err; err != nil {
		return err
	}
	r.matchRule = nameOwnerChangedMatchRule

	r.signalLoop.AddHandler(&dbusutil.SignalRule{
		Path: "/org/freedesktop/DBus",
		Name: "org.freedesktop.DBus.NameOwnerChanged",
	}, func(sig *dbus.Signal) {
		if len(sig.Body) < 3 {
			return
		}
		name, _ := sig.Body[0].(string)
		oldOwner, _ := sig.Body[1].(string)
		newOwner, _ := sig.Body[2].(string)
		if strings.HasPrefix(name, ":") && oldOwner != "" && newOwner == "" {
			r.RemoveCaller(name)
		}
	})
	return nil
}

// AddCaller registers uniqueName as an authorized caller for scope. The
// sender must be root, or a member of the privileged group that also owns the
// target process (the target must be a descendant of the sender process).
func (r *Registry) AddCaller(scope string, sender dbus.Sender, uniqueName string) error {
	if r == nil {
		return errors.New("allow-caller registry is nil")
	}
	if scope == "" {
		return errors.New("scope is empty")
	}
	if sender == "" {
		return errors.New("D-Bus sender is empty")
	}
	if !strings.HasPrefix(uniqueName, ":") {
		return fmt.Errorf("invalid D-Bus unique name %q", uniqueName)
	}

	hasOwner, err := r.service.NameHasOwner(uniqueName)
	if err != nil {
		return fmt.Errorf("check D-Bus owner %q failed: %w", uniqueName, err)
	}
	if !hasOwner {
		return fmt.Errorf("D-Bus caller %q has no owner", uniqueName)
	}
	if err := r.authorizeRegistrar(sender, uniqueName); err != nil {
		return err
	}

	r.persistMu.Lock()
	defer r.persistMu.Unlock()

	r.mu.Lock()
	callers := r.callers[scope]
	if callers == nil {
		callers = make(map[string]struct{})
		r.callers[scope] = callers
	}
	if _, exists := callers[uniqueName]; exists {
		r.mu.Unlock()
		return nil
	}
	callers[uniqueName] = struct{}{}
	if err := r.saveLocked(); err != nil {
		delete(r.callers[scope], uniqueName)
		if len(r.callers[scope]) == 0 {
			delete(r.callers, scope)
		}
		r.mu.Unlock()
		return err
	}
	r.mu.Unlock()

	logger.Infof("registered allow-caller %s for %s", uniqueName, scope)
	return nil
}

func (r *Registry) authorizeRegistrar(sender dbus.Sender, uniqueName string) error {
	senderUID, err := r.service.GetConnUID(string(sender))
	if err != nil {
		return fmt.Errorf("get SetAllowCaller sender %s UID failed: %w", sender, err)
	}
	// Root is trusted to register a process running under another account.
	if senderUID == 0 {
		return nil
	}
	if r.privilegedGroupID == invalidGroupID {
		return fmt.Errorf("privileged group %s is unavailable", privilegedGroup)
	}

	groups, err := r.service.GetConnGroups(string(sender))
	if err != nil {
		return fmt.Errorf("get SetAllowCaller sender %s groups failed: %w", sender, err)
	}
	if !containsGroup(groups, r.privilegedGroupID) {
		return fmt.Errorf("D-Bus caller %s is not in privileged group %s", sender, privilegedGroup)
	}

	targetUID, err := r.service.GetConnUID(uniqueName)
	if err != nil {
		return fmt.Errorf("get target caller %s UID failed: %w", uniqueName, err)
	}
	if targetUID != senderUID {
		return fmt.Errorf("SetAllowCaller sender UID %d does not own target %s with UID %d", senderUID, uniqueName, targetUID)
	}

	senderPID, err := r.service.GetConnPID(string(sender))
	if err != nil {
		return fmt.Errorf("get SetAllowCaller sender %s PID failed: %w", sender, err)
	}
	targetPID, err := r.service.GetConnPID(uniqueName)
	if err != nil {
		return fmt.Errorf("get target caller %s PID failed: %w", uniqueName, err)
	}
	if r.processParent == nil {
		return errors.New("process ancestry resolver is unavailable")
	}
	isDescendant, err := isProcessDescendant(targetPID, senderPID, r.processParent)
	if err != nil {
		return fmt.Errorf("verify target caller %s process ancestry failed: %w", uniqueName, err)
	}
	if !isDescendant {
		return fmt.Errorf(
			"target caller %s PID %d is not a descendant of SetAllowCaller sender %s PID %d",
			uniqueName, targetPID, sender, senderPID,
		)
	}
	return nil
}

func containsGroup(groups []uint32, target uint32) bool {
	for _, group := range groups {
		if group == target {
			return true
		}
	}
	return false
}

// Authorize reports whether sender is allowed to invoke a method protected
// by scope. Root is always allowed; otherwise the sender's unique name must
// have been registered for the scope via AddCaller.
func (r *Registry) Authorize(scope string, sender dbus.Sender) error {
	if r == nil {
		return errors.New("allow-caller registry is nil")
	}
	if sender == "" {
		return errors.New("D-Bus sender is empty")
	}

	uid, err := r.service.GetConnUID(string(sender))
	if err != nil {
		return fmt.Errorf("get caller %s UID failed: %w", sender, err)
	}
	if uid == 0 {
		return nil
	}

	r.mu.RLock()
	_, allowed := r.callers[scope][string(sender)]
	r.mu.RUnlock()
	if !allowed {
		return fmt.Errorf("D-Bus caller %s is not authorized for %s", sender, scope)
	}
	return nil
}

// RemoveCaller drops uniqueName from every scope. It is called automatically
// when the bus reports the unique name has no owner.
func (r *Registry) RemoveCaller(uniqueName string) {
	if r == nil || uniqueName == "" {
		return
	}

	r.persistMu.Lock()
	defer r.persistMu.Unlock()

	changed := false
	r.mu.Lock()
	for scope, callers := range r.callers {
		if _, exists := callers[uniqueName]; !exists {
			continue
		}
		delete(callers, uniqueName)
		changed = true
		if len(callers) == 0 {
			delete(r.callers, scope)
		}
	}
	if !changed {
		r.mu.Unlock()
		return
	}
	if err := r.saveLocked(); err != nil {
		logger.Warningf("failed to persist removal of allow-caller %s: %v", uniqueName, err)
	}
	r.mu.Unlock()
}

func (r *Registry) saveLocked() error {
	state := make(map[string][]string, len(r.callers))
	for scope, callers := range r.callers {
		uniqueNames := make([]string, 0, len(callers))
		for uniqueName := range callers {
			uniqueNames = append(uniqueNames, uniqueName)
		}
		sort.Strings(uniqueNames)
		state[scope] = uniqueNames
	}

	if r.busID == "" {
		return errors.New("system bus ID is empty")
	}
	persisted := persistedState{BusID: r.busID, Callers: state}
	if r.persistState != nil {
		return r.persistState(persisted)
	}
	return r.writeState(persisted)
}

func (r *Registry) writeState(state persistedState) error {
	content, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("marshal allow-caller state failed: %w", err)
	}

	dir := filepath.Dir(r.stateFile)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("create allow-caller runtime directory failed: %w", err)
	}
	tmp, err := ioutil.TempFile(dir, ".allow-callers-")
	if err != nil {
		return fmt.Errorf("create allow-caller state file failed: %w", err)
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)

	if err := tmp.Chmod(0600); err != nil {
		tmp.Close()
		return err
	}
	if _, err := tmp.Write(content); err != nil {
		tmp.Close()
		return fmt.Errorf("write allow-caller state failed: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("close allow-caller state failed: %w", err)
	}
	if err := os.Rename(tmpName, r.stateFile); err != nil {
		return fmt.Errorf("replace allow-caller state failed: %w", err)
	}
	return nil
}

func (r *Registry) load() error {
	content, err := ioutil.ReadFile(r.stateFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var state persistedState
	if err := json.Unmarshal(content, &state); err != nil {
		return err
	}
	if state.BusID == "" || state.BusID != r.busID {
		return nil
	}

	for scope, callers := range state.Callers {
		if scope == "" {
			continue
		}
		for _, uniqueName := range callers {
			if !strings.HasPrefix(uniqueName, ":") {
				continue
			}
			hasOwner, err := r.service.NameHasOwner(uniqueName)
			if err != nil || !hasOwner {
				continue
			}
			if r.callers[scope] == nil {
				r.callers[scope] = make(map[string]struct{})
			}
			r.callers[scope][uniqueName] = struct{}{}
		}
	}
	return nil
}
