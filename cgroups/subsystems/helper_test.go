package subsystems

import (
	"fmt"
	faker "github.com/dmgk/faker"
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestFindCgroupMountpoint(t *testing.T) {
	t.Logf("cpu subsystem mount point %v\n", FindCgroupMountpoint("cpu"))
	assert.Contains(t, FindCgroupMountpoint("cpu"), "/cpu", "cpu cgroup mount point is woring")

	t.Logf("cpuset subsystem mount point %v\n", FindCgroupMountpoint("cpuset"))
	assert.Contains(t, FindCgroupMountpoint("cpuset"), "/cpuset", "cpuset cgroup mount point is woring")

	t.Logf("memory subsystem mount point %v\n", FindCgroupMountpoint("memory"))
	assert.Contains(t, FindCgroupMountpoint("memory"), "/memory", "memory cgroup mount point is woring")
}

func TestGetCgroupPath(t *testing.T) {
	subsystem := faker.Lorem().Characters(17)
	cgroupPath := faker.Lorem().Characters(17)
	path, err := GetCgroupPath(subsystem, cgroupPath, true)
	assert.Equal(t, path, cgroupPath, "cgroup path is woring")
	assert.Nil(t, err)

	path, err = GetCgroupPath("memory", cgroupPath, true)
	assert.NotNil(t, err)
}
