package container

import (
	log "github.com/sirupsen/logrus"
	"os"
	"syscall"
	"io/ioutil"
	"strings"
)

// RunContainerInitProcess : Start container init process
func RunContainerInitProcess(command string, args []string) error {
	log.Infof("command %s", command)

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		log.Errorf(err.Error())
	}
	return nil
}

// readUserCommand : retrieve user command from command line
func readUserCommand() []string {
	pipe := os.NewFile(uintptr(3), "pipe")
	msg, err := ioutil.ReadAll(pipe)
	if err != nil {
		log.Errorf("init read pip error %")
	}
	msgStr := string(msg)
	return strings.Split(msgStr, "")
}

//// setUpMount : setup mount for system
//func setUpMount() {
//	pwd, err := os.Getwd()
//	if err != nil {
//		log.Errorf("Get current location error %v", err)
//		return
//	}
//	log.Infof("Current location is %s", pwd)
//	pivotRoot(pwd)
//
//	// mount proc
//	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
//	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
//
//	// mount tmpfs
//	defaultMountFlags = syscall.MS_NOSUID | syscall.MS_STRICTATIME
//	syscall.Mount("tmpfs", "/dev", "tmpfs", uintptr(defaultMountFlags), "mode=755" )
//}
//
//// pivotRoot : pivot root, make current working dir (pwd) as new root
//func pivotRoot(root string) error {
//	/**
//	in order t make old and new root in different fs, we have to re mount root.
//	bind mount can change mount point for same dir
//	 */
//	defaultMountFlags := syscall.MS_BIND | syscall.MS_REC
//	if err := syscall.Mount(root, root, "bind", uintptr(defaultMountFlags) , ""); err != nil {
//		return fmt.Errorf("Mount rootfs to itself error: %v", err)
//	}
//
//	// create rootfs/.pivot_root for old root
//	pivotDir := filepath.Join(root, ".pivot_root")
//	if err := os.Mkdir(pivotDir, 0777); err != nil {
//		return err
//	}
//
//	// pivot root to new rootfs, now, old_root is mount at /rootfs/.pivot_root
//	if err := syscall.PivotRoot(root, pivotDir); err != nil {
//		return fmt.Errorf("pivot_root %v", err)
//	}
//
//	// change current working dir to /
//	if err := syscall.Chdir("/"); err != nil {
//		return fmt.Errorf("chdir / %v", err)
//	}
//
//	pivotDir = filepath.Join("/", ".pivot_root")
//	// umount rootfs/.pivot_root
//	if err := syscall.Unmount(pivotDir, syscall.MNT_DETACH); err != nil {
//		return fmt.Errorf("unmount pivot_root dir %v", err)
//	}
//
//	return os.Remove(pivotDir)
//}
