package subsystems

// ResourceConfig : system resource config data type
type ResourceConfig struct {
	// memroy limitation
	MemoryLimit string
	// priority of cpu usage
	CPUShare string
	// No of cpu cores
	CPUSet string
}

// Subsystem : subsystem for container
type Subsystem interface {
	//get subsystem name, like cpu, memory
	Name() string
	// set limitation of subsystem
	Set(path string, res *ResourceConfig) error
	// apply limitaion to process
	Apply(path string, pid int) error
	// remove cgroup
	Remove(path string) error
}

var (
	// SubsystemsIns : subsystem data set
	SubsystemsIns = []Subsystem{
			&MemorySubSystem{},
	}
)
