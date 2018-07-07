package cgroups

import (
	"github.com/inputx/gocker/cgroups/subsystems"
	log "github.com/Sirupsen/logrus"
)

type CgroupManager struct {
	Path string
	Resource *subsystems.ResourceConfig
}

// NewCgroupManager : create new cgroup manager
func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{
		Path: path,
	}
}
// Apply : add pid into cgroup in order to apply limit to rocess
func (c *CgroupManager) Apply(pid int) error {
	for _, subSysIns := range(subsystems.SubsystemsIns) {
		subSysIns.Apply(c.Path, pid)
	}
	return nil
}

// Set : set subsystem limitation
func (c *CgroupManager) Set(res *subsystems.ResourceConfig) error {
	for _, subSysIns := range(subsystems.SubsystemsIns) {
		subSysIns.Set(c.Path, res)
	}
	return nil
}

// Destroy : destroy subsystem, in order to release resources
func (c *CgroupManager) Destroy() error {
	for _, subSysIns := range(subsystems.SubsystemsIns) {
		if err := subSysIns.Remove(c.Path); err != nil {
			log.Warnf("remove cgroup fail %v", err)
		}
	}
	return nil
}




