// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package cgroup

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/linuxdeepin/go-lib/log"
)

var logger = log.NewLogger("lib/cgroup")

const (
	Blkio     = "blkio"
	Cpu       = "cpu"
	Cpuacct   = "cpuacct"
	Cpuset    = "cpuset"
	Devices   = "devices"
	Freezer   = "freezer"
	Memory    = "memory"
	NetCLS    = "net_cls"
	NetPrio   = "net_prio"
	PerfEvent = "perf_event"
	Pids      = "pids"
)

var initialized bool
var mountTable []*MountTableItem

type MountTableItem struct {
	name        string
	mountPoints []string
}

func (item *MountTableItem) Name() string {
	return item.name
}

func (item *MountTableItem) MountPoints() []string {
	return item.mountPoints
}

func (item *MountTableItem) addDuplicateMount(path string) {
	item.mountPoints = append(item.mountPoints, path)
}

func Init() error {
	if initialized {
		return nil
	}
	// read /proc/cgroup
	procCgroupF, err := os.Open("/proc/cgroups")
	if err != nil {
		return err
	}

	defer procCgroupF.Close()
	procCgroupRd := bufio.NewReader(procCgroupF)

	//#subsys_name    hierarchy   num_cgroups enabled
	var subsysName string
	var hierarchy int
	var numCgroups int
	var enabled int

	// discard first line
	_, err = procCgroupRd.ReadBytes('\n')
	if err != nil {
		return err
	}
	var controllers []string
	for {
		_, err := fmt.Fscanf(procCgroupRd, "%s %d %d %d\n", &subsysName, &hierarchy, &numCgroups, &enabled)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		controllers = append(controllers, subsysName)
	}

	logger.Debug("controllers:", controllers)

	procMountsF, err := os.Open("/proc/mounts")
	if err != nil {
		return err
	}
	defer procMountsF.Close()
	procMountsRd := bufio.NewReader(procMountsF)

	for {
		mntEnt, err := getMountEntry(procMountsRd)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if mntEnt.type0 != "cgroup" {
			continue
		}

		for _, controller := range controllers {
			if mntEnt.hasOpt(controller) == "" {
				continue
			}

			logger.Debugf("found %s in %s", controller, mntEnt.opts)

			var duplicate bool
			var mntTableEntry *MountTableItem
			for _, mntTableEnt := range mountTable {
				if mntTableEnt.name == controller {
					duplicate = true
					mntTableEntry = mntTableEnt
					break
				}
			}

			if duplicate {
				logger.Debugf("controller %s is already mounted on %s",
					controller, mntTableEntry.mountPoints[0])
				mntTableEntry.addDuplicateMount(mntEnt.dir)
				continue
			}

			mountTable = append(mountTable, &MountTableItem{
				name:        controller,
				mountPoints: []string{mntEnt.dir},
			})
			logger.Debug("found cgroup option ", mntEnt.opts)
		}

		// doesn't match the controller
		// check if it is a named hierarchy.
		mntopt := mntEnt.hasOpt("name")
		if mntopt != "" {

			var duplicate bool
			var mntTableEntry *MountTableItem
			for _, mntTableEnt := range mountTable {
				if mntTableEnt.name == mntopt {
					duplicate = true
					mntTableEntry = mntTableEnt
					break
				}
			}

			if duplicate {
				logger.Debugf("controller %s is already mounted on %s",
					mntopt, mntTableEntry.mountPoints[0])
				mntTableEntry.addDuplicateMount(mntEnt.dir)
				continue
			}

			mountTable = append(mountTable, &MountTableItem{
				name:        mntopt,
				mountPoints: []string{mntEnt.dir},
			})
			logger.Debug("found cgroup option ", mntEnt.opts)
		}
	}

	//spew.Dump(mountTable)
	initialized = true
	return nil
}

func panicIfNotInitialized() {
	if !initialized {
		panic("lib cgroup is not initialized")
	}
}

func GetSubSysMountPoint(controller string) string {
	panicIfNotInitialized()
	for _, mountTableItem := range mountTable {
		if mountTableItem.name == controller {
			return mountTableItem.mountPoints[0]
		}
	}
	return ""
}

func GetSubSysMountPoints(controller string) []string {
	panicIfNotInitialized()
	for _, mountTableItem := range mountTable {
		if mountTableItem.name == controller {
			return mountTableItem.mountPoints
		}
	}
	return nil
}

type MountEntry struct {
	fsname string
	dir    string
	type0  string
	opts   string
	freq   int
	passno int
}

func (me *MountEntry) hasOpt(opt string) string {
	optEqual := opt + "="
	opts := strings.Split(me.opts, ",")
	for _, opt0 := range opts {
		if opt0 == opt || strings.HasPrefix(opt0, optEqual) {
			return opt0
		}
	}
	return ""
}

func getMountEntry(rd io.Reader) (*MountEntry, error) {
	// fsname    dir     type0  opts freq passno
	var mntent MountEntry
	_, err := fmt.Fscanf(rd, "%s %s %s %s %d %d\n", &mntent.fsname, &mntent.dir, &mntent.type0,
		&mntent.opts, &mntent.freq, &mntent.passno)
	if err != nil {
		return nil, err
	}
	return &mntent, nil
}

type Cgroup struct {
	name        string
	controllers []*Controller

	tasksUid        int
	tasksGid        int
	tasksFilePerm   os.FileMode
	controlUid      int
	controlGid      int
	controlFilePerm os.FileMode
	controlDirPerm  os.FileMode
}

func NewCgroup(name string) *Cgroup {
	cg := &Cgroup{
		name: name,
	}

	cg.init()
	return cg
}

const noUidGid = -1
const noPerms = ^os.FileMode(0)

func (cg *Cgroup) init() {
	cg.controlUid = noUidGid
	cg.controlGid = noUidGid
	cg.tasksUid = noUidGid
	cg.tasksGid = noUidGid

	cg.tasksFilePerm = noPerms
	cg.controlFilePerm = noPerms
	cg.controlDirPerm = noPerms
}

func (cg *Cgroup) Name() string {
	return cg.name
}

func (cg *Cgroup) AddController(name string) *Controller {
	ctl := &Controller{
		name:   name,
		cgroup: cg,
	}

	cg.controllers = append(cg.controllers, ctl)
	return ctl
}

func (cg *Cgroup) GetController(name string) *Controller {
	for _, ctl := range cg.controllers {
		if ctl.name == name {
			return ctl
		}
	}
	return nil
}

func (cg *Cgroup) NewChildGroup(name string) *Cgroup {
	child := NewCgroup(filepath.Join(cg.name, name))
	child.SetUidGid(cg.GetUidGid())
	child.SetPermissions(cg.controlDirPerm, cg.controlFilePerm, cg.tasksFilePerm)
	for _, ctl := range cg.controllers {
		child.AddController(ctl.name)
	}
	return child
}

func (cg *Cgroup) NewChildrenGroups(controller string) ([]*Cgroup, error) {
	childrenNames, err := cg.GetChildren(controller)
	if err != nil {
		return nil, err
	}
	result := make([]*Cgroup, len(childrenNames))
	for idx, name := range childrenNames {
		result[idx] = cg.NewChildGroup(name)
	}
	return result, nil
}

func (cg *Cgroup) GetChildren(controller string) ([]string, error) {
	panicIfNotInitialized()

	if controller == "" {
		controller = cg.getFirstControllerName()
	}
	dir := buildPath(cg.name, controller)

	fileInfoList, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, fileInfo := range fileInfoList {
		if fileInfo.IsDir() {
			result = append(result, fileInfo.Name())
		}
	}
	return result, nil
}

func testSubSysMounted(name string) bool {
	for _, mountTableItem := range mountTable {
		if mountTableItem.name == name {
			return true
		}
	}
	return false
}

var ErrSubSysNotMounted = errors.New("cgroup one of the needed subsystems is not mounted")

// Create physically create a control group in kernel.
func (cg *Cgroup) Create(ignoreOwnership bool) (err error) {
	panicIfNotInitialized()

	for _, ctl := range cg.controllers {
		if !testSubSysMounted(ctl.name) {
			return ErrSubSysNotMounted
		}
	}

	for _, ctl := range cg.controllers {
		path := buildPath(cg.name, ctl.name)
		base := path
		if path == "" {
			continue
		}

		err = createControlGroup(path)
		if err != nil {
			return err
		}

		if !ignoreOwnership {
			logger.Debug("changing ownership of ", path)
			err = chownRecursive(path, cg.controlUid, cg.controlGid)
			if err == nil {
				err = chmodRecursiveController(path,
					cg.controlDirPerm,
					cg.controlDirPerm != noPerms,
					cg.controlFilePerm,
					cg.controlFilePerm != noPerms,
					true,
					[]string{"tasks"})
			}
		}

		if err != nil {
			return err
		}

		if !ignoreOwnership {
			path := filepath.Join(base, "tasks")
			err = chown(path, cg.tasksUid, cg.tasksGid)
			if err == nil && cg.tasksFilePerm != noPerms {
				var info os.FileInfo
				info, err = os.Stat(path)
				if err == nil {
					err = chmodPath(path, info, cg.tasksFilePerm, true)
				}
			}

			if err != nil {
				return err
			}
		}
	}
	return nil
}

var ErrCgroupNotMounted = errors.New("cgroup is not mounted")

func createControlGroup(path string) error {
	if !testMountedFs() {
		return ErrCgroupNotMounted
	}
	return os.MkdirAll(path, 0755)
}

func chownRecursive(path string, owner, group int) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		err = chown(path, owner, group)
		if err != nil {
			logger.Warningf("can not change owner of file %s : %v", path, err)
		}

		return nil
	})
}

func chmodRecursiveController(path string,
	dirMode os.FileMode, dirModeChange bool,
	fileMode os.FileMode, fileModeChange bool,
	ownerIsUmask bool, ignoreList []string) error {
	logger.Debug("chmod: path is ", path)

	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, ig := range ignoreList {
			if ig == path {
				// ignore
				return nil
			}
		}
		err = chmodFile(path, info, dirMode, dirModeChange, fileMode, fileModeChange, ownerIsUmask)
		if err != nil {
			logger.Warningf("can not change file mode %s: %v", path, err)
		}
		return nil
	})
}

func chown(filename string, owner, group int) error {
	if owner == noUidGid {
		owner = os.Getuid()
	}

	if group == noUidGid {
		group = os.Getgid()
	}

	return os.Chown(filename, owner, group)
}

func chmodFile(filename string, info os.FileInfo,
	dirMode os.FileMode, dirModeChange bool,
	fileMode os.FileMode, fileModeChange bool,
	ownerIsUmask bool) error {

	if info.IsDir() {
		if dirModeChange {
			return chmodPath(filename, info, dirMode, ownerIsUmask)
		}
	} else {
		if fileModeChange {
			return chmodPath(filename, info, fileMode, ownerIsUmask)
		}
	}

	return nil
}

func chmodPath(path string, info os.FileInfo, mode os.FileMode, ownerIsUmask bool) error {
	mask := ^os.FileMode(0)
	if ownerIsUmask {
		var umask, gmask, omask os.FileMode

		umask = syscall.S_IRWXU & info.Mode()
		gmask = umask >> 3
		omask = umask >> 3

		mask = umask | gmask | omask | syscall.S_ISUID | syscall.S_ISGID | syscall.S_ISVTX
	}

	return os.Chmod(path, mode&mask)
}

func testMountedFs() bool {
	procMountsF, err := os.Open("/proc/mounts")
	if err != nil {
		return false
	}
	defer procMountsF.Close()
	procMountsRd := bufio.NewReader(procMountsF)

	for {
		mntEnt, err := getMountEntry(procMountsRd)
		if err != nil {
			return false
		}

		if mntEnt.type0 == "cgroup" {
			return true
		}
	}
}

func buildPath(name, controller string) string {
	for _, mountTableItem := range mountTable {
		if mountTableItem.name == controller {
			return filepath.Join(mountTableItem.mountPoints[0], name)
		}
	}
	return ""
}

type DeleteFlags uint

const (
	DeleteFlagIgnoreMigration = 1 << iota
	DeleteFlagRecursive
	DeleteFlagEmptyOnly
)

var ErrInvalid = errors.New("invalid operation")

func (cg *Cgroup) Delete(flags DeleteFlags) (err error) {
	panicIfNotInitialized()

	deleteCgroup := true
	if (flags&DeleteFlagEmptyOnly != 0) &&
		(flags&DeleteFlagRecursive != 0) {
		return ErrInvalid
	}

	for _, ctl := range cg.controllers {
		if !testSubSysMounted(ctl.name) {
			return ErrSubSysNotMounted
		}
	}

	var firstErr error
	var parentName string
	for _, ctl := range cg.controllers {
		if flags&DeleteFlagEmptyOnly == 0 {
			parentName, err = cg.findParent(ctl.name)
			if err != nil {
				continue
			}

			if parentName == "" {
				// root cgroup is being deleted
				if flags&DeleteFlagRecursive != 0 {
					// move all tasks to the root group and
					// do not delete it afterwards.
					parentName = "."
					deleteCgroup = false
				} else {
					continue
				}
			}
		}

		var parentTasksF *os.File

		logger.Debug("parentName is", parentName)
		if parentName != "" {
			// tasks need to be moved, pre-open target tasks file
			parentPath := buildPath(parentName, ctl.name)
			parentTasksPath := filepath.Join(parentPath, "tasks")
			parentTasksF, err = os.OpenFile(parentTasksPath, os.O_WRONLY, 0644)
			if err != nil {
				logger.Debugf("warnning: can not open file %q: %v", parentTasksPath, err)
				continue
			}
		}

		if flags&DeleteFlagRecursive != 0 {
			err = deleteCgroupControllerRecursive(cg.name, ctl.name, parentTasksF, flags, deleteCgroup)
		} else {
			err = deleteCgroupController(cg.name, ctl.name, parentTasksF, flags)
		}

		if parentTasksF != nil {
			parentTasksF.Close()
		}

		// ErrNonEmpty is more or less not an error,
		// but an indication that something was not removed.
		// Therefore it should be replaced by any other error.
		if err != nil {
			if firstErr == nil {
				firstErr = err
			} else if firstErr == ErrNonEmpty && firstErr != err {
				firstErr = err
			}
			logger.Warningf("failed to remove group %q: %v", ctl.name+":"+cg.name, err)
		}
	}

	return firstErr
}

var ErrNonEmpty = errors.New("group is not empty")

func deleteCgroupController(cgroupName, controller string, targetTasksF *os.File, flags DeleteFlags) (err error) {
	logger.Debugf("Removing group %s:%s", controller, cgroupName)

	if flags&DeleteFlagEmptyOnly == 0 {
		// open tasks file of the group to delete
		path := filepath.Join(buildPath(cgroupName, controller), "tasks")
		logger.Debugf("tasks path is %q", path)
		var deleteTaskF *os.File
		deleteTaskF, err = os.Open(path)
		if err == nil {
			err = moveTaskFiles(deleteTaskF, targetTasksF)
			if err != nil {
				logger.Warning("moveTaskFiles err:", err)
			}
			deleteTaskF.Close()
		} else {
			// can't open the tasks file. if the file does not
			// exist, ignore it - the group has been already removed.
			if os.IsNotExist(err) {
				err = nil
			} else {
				logger.Warningf("can not open %q: %v", path, err)
			}
		}

		if err != nil && flags&DeleteFlagIgnoreMigration == 0 {
			return err
		}
	}
	// remove the group
	path := buildPath(cgroupName, controller)
	err = os.Remove(path)
	if err == nil || os.IsNotExist(err) {
		return nil
	}

	if flags&DeleteFlagEmptyOnly != 0 {
		if pErr, ok := err.(*os.PathError); ok && pErr.Err == syscall.EBUSY {
			return ErrNonEmpty
		}
	}

	logger.Warningf("can not remove directory %q:%v", path, err)
	return err
}

func moveTaskFiles(inputTasksF, outputTasksF *os.File) error {
	logger.Debug("move tasks")
	var tid int
	for {
		_, err := fmt.Fscanln(inputTasksF, &tid)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		logger.Debug("moveTaskFiles tid:", tid)

		_, err = fmt.Fprintln(outputTasksF, tid)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteCgroupControllerRecursive(cgroupName, controller string, targetTasksF *os.File,
	flags DeleteFlags, deleteRoot bool) error {

	logger.Debugf("Recursively removing %s:%s", controller, cgroupName)

	path := buildPath(cgroupName, controller)

	fd, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	for {
		names, err1 := fd.Readdirnames(100)
		for _, name := range names {
			cgroupName1 := cgroupName + string(os.PathSeparator) + name

			p := buildPath(cgroupName1, controller)
			pStat, err := os.Stat(p)
			if err != nil {
				continue
			}

			if pStat.IsDir() {
				_ = deleteCgroupControllerRecursive(cgroupName1, controller, targetTasksF, flags, true)
			}

		}
		if err1 == io.EOF {
			break
		}

		if len(names) == 0 {
			break
		}
	}
	_ = fd.Close()

	if deleteRoot {
		err = deleteCgroupController(cgroupName, controller, targetTasksF, flags)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cg *Cgroup) getParentName() string {
	dir := cg.name
	pdir := filepath.Dir(dir)

	// check for root cgroup
	if cg.name == "" || cg.name == pdir {
		logger.Debugf("specified cgroup %q is root cgroup", cg.name)
		return ""
	}
	return pdir
}

func (cg *Cgroup) findParent(controller string) (string, error) {
	childPath := buildPath(cg.name, controller)
	parentPath := filepath.Join(childPath, "..")

	logger.Debugf("findParent child path: %q", childPath)
	logger.Debugf("findParent parent path: %q", parentPath)
	statChild, err := os.Stat(childPath)
	if err != nil {
		return "", err
	}
	statParent, err := os.Stat(parentPath)
	if err != nil {
		return "", err
	}

	childDev := statChild.Sys().(*syscall.Stat_t).Dev
	parentDev := statParent.Sys().(*syscall.Stat_t).Dev
	if childDev != parentDev {
		logger.Debug("parent is on different device")
		return "", nil
	}
	return cg.getParentName(), nil
}

func (cg *Cgroup) AllExist() bool {
	panicIfNotInitialized()

	for _, ctl := range cg.controllers {
		if !ctl.exists() {
			return false
		}
	}
	return true
}

func (cg *Cgroup) SetUidGid(tasksUid, tasksGid, controlUid, controlGid int) {
	cg.tasksUid = tasksUid
	cg.tasksGid = tasksGid
	cg.controlUid = controlUid
	cg.controlGid = controlGid
}

func (cg *Cgroup) GetUidGid() (int, int, int, int) {
	return cg.tasksUid, cg.tasksGid, cg.controlUid, cg.controlGid
}

func (cg *Cgroup) SetPermissions(controlDirPerm, controlFilePerm, tasksFilePerm os.FileMode) {
	cg.controlDirPerm = controlDirPerm
	cg.controlFilePerm = controlFilePerm
	cg.tasksFilePerm = tasksFilePerm
}

func (cg *Cgroup) AttachCurrentTask() error {
	return cg.AttachTask(syscall.Gettid())
}

func (cg *Cgroup) AttachTask(tid int) error {
	panicIfNotInitialized()

	for _, ctl := range cg.controllers {
		if !testSubSysMounted(ctl.name) {
			return ErrSubSysNotMounted
		}
	}

	for _, ctl := range cg.controllers {
		path := buildPath(cg.name, ctl.name)
		taskPath := filepath.Join(path, "tasks")
		err := attachTaskProc(taskPath, tid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cg *Cgroup) AttachCurrentProcess() error {
	return cg.AttachProcess(os.Getpid())
}

func (cg *Cgroup) AttachProcess(pid int) error {
	panicIfNotInitialized()

	for _, ctl := range cg.controllers {
		if !testSubSysMounted(ctl.name) {
			return ErrSubSysNotMounted
		}
	}

	for _, ctl := range cg.controllers {
		cgroupProcsFile := filepath.Join(buildPath(cg.name, ctl.name), "cgroup.procs")
		err := attachTaskProc(cgroupProcsFile, pid)
		if err != nil {
			return err
		}
	}
	return nil
}

func attachTaskProc(path string, tid int) error {
	taskF, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer taskF.Close()
	_, err = fmt.Fprintln(taskF, tid)
	if err != nil {
		return err
	}

	return taskF.Sync()
}

func getIntegers(name, controller, basename string) ([]int, error) {
	panicIfNotInitialized()

	filename := filepath.Join(buildPath(name, controller), basename)
	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	var ints []int
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		pid, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		ints = append(ints, pid)
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return ints, nil
}

func GetTasks(name, controller string) ([]int, error) {
	return getIntegers(name, controller, "tasks")
}

func GetProcs(name, controller string) ([]int, error) {
	return getIntegers(name, controller, "cgroup.procs")
}

func (cg *Cgroup) GetProcs(controller string) ([]int, error) {
	if controller == "" {
		controller = cg.getFirstControllerName()
		if controller == "" {
			return nil, errors.New("no controller")
		}
	}
	return GetProcs(cg.name, controller)
}

func (cg *Cgroup) GetTasks(controller string) ([]int, error) {
	if controller == "" {
		controller = cg.getFirstControllerName()
		if controller == "" {
			return nil, errors.New("no controller")
		}
	}
	return GetTasks(cg.name, controller)
}

func (cg *Cgroup) getFirstControllerName() string {
	for _, ctl := range cg.controllers {
		return ctl.name
	}
	return ""
}

func GetProcessControllerPath(pid int, controller string) (string, error) {
	procCgroupFile := fmt.Sprintf("/proc/%d/cgroup", pid)
	pidCgroupF, err := os.Open(procCgroupFile)
	if err != nil {
		return "", err
	}
	defer pidCgroupF.Close()

	scanner := bufio.NewScanner(pidCgroupF)
	for scanner.Scan() {
		line := scanner.Bytes()
		parts := bytes.SplitN(line, []byte{':'}, 3)
		if len(parts) != 3 {
			continue
		}
		ctlName := string(parts[1])
		cgroupPath := string(parts[2])
		if ctlName == controller {
			return cgroupPath, nil
		}
	}
	err = scanner.Err()
	if err != nil {
		return "", err
	}
	return "", errors.New("controller not found")
}

type Controller struct {
	name   string
	cgroup *Cgroup
}

func (c *Controller) Name() string {
	return c.name
}

func (c *Controller) Cgroup() *Cgroup {
	return c.cgroup
}

func (c *Controller) Exists() bool {
	panicIfNotInitialized()
	return c.exists()
}

func (c *Controller) exists() bool {
	path := buildPath(c.cgroup.name, c.name)
	fileInfo, err := os.Stat(path)

	if err != nil || !fileInfo.IsDir() {
		return false
	}

	cgroupProcsFile := filepath.Join(path, "cgroup.procs")
	cgroupProcsFi, err := os.Stat(cgroupProcsFile)
	if err != nil || cgroupProcsFi.IsDir() {
		return false
	}

	return true
}

func (c *Controller) getValueFile(name string) string {
	return filepath.Join(buildPath(c.cgroup.name, c.name), c.name+"."+name)
}

func (c *Controller) GetValueString(name string) (string, error) {
	panicIfNotInitialized()

	filename := c.getValueFile(name)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (c *Controller) SetValueString(name, value string) error {
	panicIfNotInitialized()

	filename := c.getValueFile(name)
	fh, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fh.Close()
	_, err = fh.WriteString(value)
	if err != nil {
		return err
	}
	return fh.Sync()
}

func (c *Controller) GetValueInt64(name string) (int64, error) {
	valStr, err := c.GetValueString(name)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(strings.TrimSpace(valStr), 10, 64)
}

func (c *Controller) SetValueInt64(name string, value int64) error {
	valStr := strconv.FormatInt(value, 10)
	return c.SetValueString(name, valStr)
}

func (c *Controller) GetValueUint64(name string) (uint64, error) {
	valStr, err := c.GetValueString(name)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(strings.TrimSpace(valStr), 10, 64)
}

func (c *Controller) SetValueUint64(name string, value uint64) error {
	valStr := strconv.FormatUint(value, 10)
	return c.SetValueString(name, valStr)
}

func (c *Controller) GetValueBool(name string) (bool, error) {
	valStr, err := c.GetValueString(name)
	if err != nil {
		return false, err
	}
	valInt, err := strconv.ParseInt(strings.TrimSpace(valStr), 10, 64)
	if err != nil {
		return false, err
	}

	return valInt != 0, nil
}

func (c *Controller) SetValueBool(name string, value bool) error {
	valStr := "0"
	if value {
		valStr = "1"
	}
	return c.SetValueString(name, valStr)
}

func (c *Controller) GetStats(fields []string, args ...*uint64) (int, error) {
	panicIfNotInitialized()

	if len(fields) != len(args) {
		return 0, errors.New("len(fields) != len(args)")
	}

	statFile := c.getValueFile("stat")
	statF, err := os.Open(statFile)
	if err != nil {
		return 0, err
	}
	defer statF.Close()

	scanner := bufio.NewScanner(statF)
	var count int
	for scanner.Scan() {
		parts := bytes.Fields(scanner.Bytes())
		if len(parts) != 2 {
			continue
		}
		key := string(parts[0])

		for i := 0; i < len(fields); i++ {
			if fields[i] == key {
				valUint, err := strconv.ParseUint(string(parts[1]), 10, 64)
				if err != nil {
					return count, err
				}
				*(args[i]) = valUint
				count++
				break
			}
		}

		if count == len(fields) {
			break
		}
	}
	return count, scanner.Err()
}

func (c *Controller) GetStatsAll() (map[string]uint64, error) {
	panicIfNotInitialized()
	statFile := c.getValueFile("stat")
	statF, err := os.Open(statFile)
	if err != nil {
		return nil, err
	}
	defer statF.Close()

	ret := make(map[string]uint64)
	scanner := bufio.NewScanner(statF)
	for scanner.Scan() {
		parts := bytes.Fields(scanner.Bytes())
		if len(parts) != 2 {
			continue
		}
		valUint, err := strconv.ParseUint(string(parts[1]), 10, 64)
		if err != nil {
			return nil, err
		}
		ret[string(parts[0])] = valUint
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type ControllerInfo struct {
	name       string
	hierarchy  int
	numCgroups int
	enabled    bool
}

func (cd *ControllerInfo) Name() string {
	return cd.name
}

// Hierarchy is the identification of the controller. Controllers with the same
// hierarchy ID are mounted together as one hierarchy. Controllers with ID 0
// are not currently mounted anywhere.
func (cd *ControllerInfo) Hierarchy() int {
	return cd.hierarchy
}

func (cd *ControllerInfo) NumCgroups() int {
	return cd.numCgroups
}

func (cd *ControllerInfo) Enabled() bool {
	return cd.enabled
}

// GetAllControllers list all controllers, including those which are not mounted.
func GetAllControllers() ([]*ControllerInfo, error) {
	procCgroupF, err := os.Open("/proc/cgroups")
	if err != nil {
		return nil, err
	}

	defer procCgroupF.Close()
	procCgroupRd := bufio.NewReader(procCgroupF)

	//#subsys_name    hierarchy   num_cgroups enabled
	var name string
	var hierarchy int
	var numCgroups int
	var enabled int

	// discard first line
	_, err = procCgroupRd.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	var controllers []*ControllerInfo
	for {
		_, err := fmt.Fscanf(procCgroupRd, "%s %d %d %d\n",
			&name, &hierarchy, &numCgroups, &enabled)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		var enabledBool bool
		if enabled != 0 {
			enabledBool = true
		}

		controllers = append(controllers, &ControllerInfo{
			name:       name,
			hierarchy:  hierarchy,
			numCgroups: numCgroups,
			enabled:    enabledBool,
		})
	}

	return controllers, nil
}

// GetControllers list mounted controllers.
func GetControllers() []*MountTableItem {
	panicIfNotInitialized()
	return mountTable
}
