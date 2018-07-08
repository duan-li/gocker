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

//func TestPivotRoot(t *testing.T) {
//	pwd, err := os.Getwd()
//	assert.Nil(t, err, "err should be nil")
//	error := pivotRoot(pwd)
//	assert.NotNil(t, error, "error should be nil")
//}

func TestReadUserCommand(t *testing.T) {
	message := readUserCommand()
	assert.NotNil(t, message, "error should be nil")
}