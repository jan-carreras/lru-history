package view

import (
	"errors"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jan-carreras/lru-history/internal/models"
)

var ErrNoCommandSelected = errors.New("no command selected")

type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Render(cmds []models.Counter) (*models.HistoryLine, error) {
	m := initialModel(cmds)

	p := tea.NewProgram(&m)
	if _, err := p.Run(); err != nil {
		return nil, err
	}

	if m.cursor >= 0 {
		return &cmds[m.cursor].Command, nil
	}

	return nil, ErrNoCommandSelected
}

func initialModel(cmds []models.Counter) model {
	ti := textinput.New()
	ti.Placeholder, ti.CharLimit, ti.Width = "Search: ", 100, 25
	ti.Focus()

	return model{
		choices:   cmds,
		textInput: ti,
	}
}
