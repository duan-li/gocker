package container

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestRunContainerInitProcess(t *testing.T) {
	command := "fake-command"
	args := []string{}
	init := RunContainerInitProcess(command, args)

	assert.Nil(t, init, "should be nil")
}

func TestReadUserCommand(t *testing.T) {
	message := readUserCommand()
	assert.NotNil(t, message, "error should be nil")
}