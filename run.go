package main

import (
	"github.com/inputx/gocker/container"
	log "github.com/sirupsen/logrus"
	"os"
)

// Run : Ran parent process
func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
