package container

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"syscall"
)

// NewParentProcess : Create a paranet process
func NewParentProcess(tty bool, command string) (*exec.Cmd, *os.File) {
	// get system pip
	readPipe, writePipe, err := Pipe()
	if err != nil {
		log.Errorf("New pipe error %v", err)
		return nil, nil
	}

	// initCommand RunContainerInitProcess
	cmd := exec.Command("/proc/self/exe", "init")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	cmd.ExtraFiles = []*os.File{readPipe}
	return cmd, writePipe
}

// Pipe : system pipe
func Pipe() (*os.File, *os.File, error) {
	read, write, err := os.Pipe()
	if err == nil {
		return read, write, nil
	}
	return nil, nil, err
}
