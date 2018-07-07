package cgroups

import (
	faker "github.com/dmgk/faker"
	assert "github.com/stretchr/testify/assert"
	"github.com/inputx/gocker/cgroups/subsystems"
	"testing"
	mock "github.com/stretchr/testify/mock"
)

type MockSubSystem struct {
	mock.Mock
}

func (m *MockSubSystem) Apply(cgroupPath string, pid int) error {
	return nil
}

func TestNewCgroupManager(t *testing.T) {
	path := faker.Lorem().Characters(17)
	manager := NewCgroupManager(path)
	assert.Equal(t, path, manager.Path)
}

func TestApply(t *testing.T) {
	path := faker.Lorem().Characters(17)

	manager := NewCgroupManager(path)
	assert.Nil(t, manager.Apply(1))
}


func TestSet(t *testing.T) {
	path := faker.Lorem().Characters(17)
	manager := NewCgroupManager(path)
	resConf := &subsystems.ResourceConfig{
		MemoryLimit: "10m",
	}
	assert.Nil(t, manager.Set(resConf))
}


func TestDestroy(t *testing.T) {
	path := faker.Lorem().Characters(17)
	manager := NewCgroupManager(path)
	assert.Nil(t, manager.Destroy())
}