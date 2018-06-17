package container

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestRunContainerInitProcess(t *testing.T) {
	command := "/bin/sh"
	args := []string{}
	init := RunContainerInitProcess(command, args)

	assert.Nil(t, init, "should be nil")
}
