package subsystems

type ResourceConfig struct {
	// memroy limitation
	MemoryLimit string
	// priority of cpu usage
	CpuShare string
	// No of cpu cores
	CpuSet string
}

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
	SubsystemsIns = []Subsystem{
	// 		&MemorySubSystem{},
	}
)
