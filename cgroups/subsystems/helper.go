package subsystems

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

// FindCgroupMountpoint : find cgroup mount point
func FindCgroupMountpoint(subsystem string) string {
	f, err := os.Open("/proc/self/mountinfo")
	if err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			txt := scanner.Text()
			fields := strings.Split(txt, " ")
			for _, opt := range strings.Split(fields[len(fields)-1], ",") {
				if opt == subsystem {
					return fields[4]
				}
			}
		}

	}

	return ""
}

// GetCgroupPath : get cgroup path
func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgroupMountpoint(subsystem)
	_, err := os.Stat(path.Join(cgroupRoot, cgroupPath))
	if err == nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path.Join(cgroupRoot, cgroupPath), 0755); err == nil {
			} else {
				return "", fmt.Errorf("error create cgroup %v", err)
			}
		}
		return path.Join(cgroupRoot, cgroupPath), nil
	}
	return "", fmt.Errorf("cgroup path error %v", err)
}
