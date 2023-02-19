package run

import (
	"github.com/jan-carreras/lru-history/internal/models"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type Runner struct{}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Run(command models.HistoryLine) error {
	cmdArgs := strings.Split(command.Command, " ")
	binary, lookErr := exec.LookPath(cmdArgs[0])
	if lookErr != nil {
		return lookErr
	}

	return syscall.Exec(binary, cmdArgs, os.Environ())
}
