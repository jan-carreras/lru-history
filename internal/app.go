// Package internal defines the app that orchestrates the entire flow of the app
package internal

import (
	"errors"
	"io"
	"sort"
	"strings"

	"github.com/jan-carreras/lru-history/internal/models"
	"github.com/jan-carreras/lru-history/internal/storage"
	"github.com/jan-carreras/lru-history/internal/view"
)

const (
	maxItems = 10
)

type renderer interface {
	Render(cmds []models.Counter) (*models.HistoryLine, error)
}

type runner interface {
	Run(command models.HistoryLine) error
}

// App orchestrate all the flows of the application
type App struct {
	storage  *storage.Storage
	renderer renderer
	runner   runner
}

// NewApp returns an App
func NewApp(storage *storage.Storage, renderer renderer, runner runner) *App {
	return &App{
		storage:  storage,
		renderer: renderer,
		runner:   runner,
	}
}

// AddToHistory adds a new command into the History file
func (a *App) AddToHistory(input io.Reader) error {
	return a.storage.AddHistoryLine(input)
}

// History allows to select a command to execute from a list
func (a *App) History(fromDir string) error {
	lines, err := a.storage.ReadHistory()
	if err != nil {
		return err
	}

	counters := make(map[string]models.Counter)
	for _, cmd := range lines {
		if !strings.HasPrefix(cmd.Directory, fromDir) {
			continue
		} else if c, ok := counters[cmd.Command]; !ok {
			counters[cmd.Command] = models.Counter{Count: 1, Command: cmd}
		} else {
			c.Count++
			if cmd.ExecutedAt.After(c.Command.ExecutedAt) {
				c.Command = cmd
			}
			counters[cmd.Command] = c
		}
	}

	cmds := make([]models.Counter, 0, len(counters))
	for _, counter := range counters {
		cmds = append(cmds, counter)
	}

	sort.Slice(cmds, func(i, j int) bool {
		return cmds[i].Count > cmds[j].Count
	})

	if len(cmds) > maxItems {
		cmds = cmds[:maxItems]
	}

	command, err := a.renderer.Render(cmds)
	if errors.Is(err, view.ErrNoCommandSelected) {
		return nil
	} else if err != nil {
		return err
	}

	return a.runner.Run(*command)
}
