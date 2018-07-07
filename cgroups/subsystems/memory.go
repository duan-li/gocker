package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

// MemorySubSystem : memroy subsystem data type
type MemorySubSystem struct {
}

// Set : Set limitation of memory subsystem
func (s *MemorySubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	// get memory cgrouppath
	subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true)
	if err == nil {
		if res.MemoryLimit != "" {
			// write config value into cgroup `memory.limit_in_bytes` file
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644); err != nil {
				return fmt.Errorf("set cgroup memory fail %v", err)
			}
		}
		return nil
	}

	return err
}

// Remove : Remove limitation of memory subsystem
func (s *MemorySubSystem) Remove(cgroupPath string) error {
	subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false)
	if err == nil {
		return os.RemoveAll(subsysCgroupPath)
	}
	return err
}

// Apply : Apply limitation of memory subsystem to process
func (s *MemorySubSystem) Apply(cgroupPath string, pid int) error {
	subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false)
	if err == nil {
		if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("set cgroup proc fail %v", err)
		}
		return nil
	}
	return fmt.Errorf("get cgroup %s error: %v", cgroupPath, err)
}

// Name : get subsystem name
func (s *MemorySubSystem) Name() string {
	return "memory"
}
