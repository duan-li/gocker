package container

import (
	faker "github.com/dmgk/faker"
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestRunContainerInitProcess(t *testing.T) {
	command := faker.Lorem().Characters(17)
	args := []string{}
	init := RunContainerInitProcess(command, args)

	assert.Nil(t, init, "should be nil")
}
