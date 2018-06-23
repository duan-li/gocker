package container

import (
	faker "github.com/dmgk/faker"
	assert "github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewParentProcessSysProcAttr(t *testing.T) {
	command := faker.Lorem().Characters(17)
	cmd, pip := NewParentProcess(true, command)

	expected := []string{"/proc/self/exe", "init"}

	assert.Equal(t, cmd.Args, expected, "cmd.Args is woring")
	assert.NotNil(t, cmd.SysProcAttr, "cmd.SysProcAttr not set")
	assert.Equal(t, cmd.Path, "/proc/self/exe", "cmd.Path is woring")

	assert.NotNil(t, pip, "pip is woring")
}

func TestNewParentProcessTtyTrue(t *testing.T) {
	cmd, pip := NewParentProcess(true, faker.Lorem().Characters(17))

	assert.Equal(t, cmd.Stdin, os.Stdin, "cmd tty error (Stdin)")
	assert.Equal(t, cmd.Stdout, os.Stdout, "cmd tty error (Stdout)")
	assert.Equal(t, cmd.Stderr, os.Stderr, "cmd tty error (Stderr)")

	assert.NotNil(t, pip, "pip is woring")
}

func TestNewParentProcessTtyFalse(t *testing.T) {
	cmd, pip := NewParentProcess(false, faker.Lorem().Characters(17))

	assert.Nil(t, cmd.Stdin, "cmd tty error (Stdin)")
	assert.Nil(t, cmd.Stdout, "cmd tty error (Stdout)")
	assert.Nil(t, cmd.Stderr, "cmd tty error (Stderr)")

	assert.NotNil(t, pip, "pip is woring")
}

func TestPipe(t *testing.T) {
	testRead, testWrite, testErr := Pipe()
	assert.NotNil(t, testRead, "read pipe is wrong")
	assert.NotNil(t, testWrite, "wrint pipe is wrong")
	assert.Nil(t, testErr, "err pipe is wrong")
}
