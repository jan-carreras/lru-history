// Package run executes a given command by running a syscall Exec
package run

import (
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/jan-carreras/lru-history/internal/models"
)

// Runner runs commands on the os
type Runner struct{}

// NewRunner returns a Runner
func NewRunner() *Runner {
	return &Runner{}
}

// Run tells the OS to Exec the given command
func (r *Runner) Run(command models.HistoryLine) error {
	cmdArgs := strings.Split(command.Command, " ")
	binary, lookErr := exec.LookPath(cmdArgs[0])
	if lookErr != nil {
		return lookErr
	}

	return syscall.Exec(binary, cmdArgs, os.Environ())
}
