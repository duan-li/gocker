package subsystems

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryName(t *testing.T) {
	memSubSys := MemorySubSystem{}
	assert.Equal(t, memSubSys.Name(), "memory", "Name is woring")
}

// func TestMemoryCgroup(t *testing.T) {
// 	memSubSys := MemorySubSystem{}
// 	resConfig := ResourceConfig{
// 		MemoryLimit: "1000m",
// 	}
// 	testCgroup := "testmemlimit"

// 	if err := memSubSys.Set(testCgroup, &resConfig); err != nil {
// 		t.Fatalf("cgroup fail %v", err)
// 	}
// 	stat, _ := os.Stat(path.Join(FindCgroupMountpoint("memory"), testCgroup))
// 	t.Logf("cgroup stats: %+v", stat)

// 	if err := memSubSys.Apply(testCgroup, os.Getpid()); err != nil {
// 		t.Fatalf("cgroup Apply %v", err)
// 	}
// 	//move process to root of Cgroup node
// 	if err := memSubSys.Apply("", os.Getpid()); err != nil {
// 		t.Fatalf("cgroup Apply %v", err)
// 	}

// 	if err := memSubSys.Remove(testCgroup); err != nil {
// 		t.Fatalf("cgroup remove %v", err)
// 	}
// }
