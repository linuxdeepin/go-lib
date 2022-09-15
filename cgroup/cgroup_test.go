// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package cgroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mountTableTest = []*MountTableItem{
	{
		name:        "name=systemd",
		mountPoints: []string{"/sys/fs/cgroup/systemd"},
	},
	{
		name:        "hugetlb",
		mountPoints: []string{"/sys/fs/cgroup/hugetlb"},
	},
	{
		name:        "freezer",
		mountPoints: []string{"/sys/fs/cgroup/freezer"},
	},
	{
		name:        "memory",
		mountPoints: []string{"/sys/fs/cgroup/memory"},
	},
	{
		name:        "net_cls",
		mountPoints: []string{"/sys/fs/cgroup/net_cls,net_prio"},
	},
	{
		name:        "net_prio",
		mountPoints: []string{"/sys/fs/cgroup/net_cls,net_prio"},
	},
	{
		name:        "cpu",
		mountPoints: []string{"/sys/fs/cgroup/cpu,cpuacct"},
	},
	{
		name:        "cpuacct",
		mountPoints: []string{"/sys/fs/cgroup/cpu,cpuacct"},
	},
	{
		name:        "cpuset",
		mountPoints: []string{"/sys/fs/cgroup/cpuset"},
	},
	{
		name:        "rdma",
		mountPoints: []string{"/sys/fs/cgroup/rdma"},
	},
	{
		name:        "blkio",
		mountPoints: []string{"/sys/fs/cgroup/blkio"},
	},
	{
		name:        "devices",
		mountPoints: []string{"/sys/fs/cgroup/devices"},
	}, {
		name:        "perf_event",
		mountPoints: []string{"/sys/fs/cgroup/perf_event"},
	},
	{
		name:        "pids",
		mountPoints: []string{"/sys/fs/cgroup/pids"},
	},
}

func Test_MountTableItem(t *testing.T) {
	err := Init()
	require.NoError(t, err)
	for _, mountTableItem := range mountTable {
		for _, mountTest := range mountTableTest {
			if mountTableItem.Name() == mountTest.name {
				assert.Equal(t, mountTableItem.MountPoints(), mountTest.mountPoints)
				assert.Equal(t, GetSubSysMountPoint(mountTableItem.name), mountTest.mountPoints[0])
				assert.Equal(t, GetSubSysMountPoints(mountTableItem.name), mountTest.mountPoints)
			}
		}
	}
}
