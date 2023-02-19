package view

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jan-carreras/lru-history/internal/models"
)

type model struct {
	choices   []models.Counter
	cursor    int
	textInput textinput.Model
}

func (m *model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m *model) Update(msg tea.Msg) (md tea.Model, cmd tea.Cmd) {
	m.textInput, cmd = m.textInput.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.cursor = -1
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			return m, tea.Quit
		default:
			n, err := strconv.Atoi(m.textInput.Value())
			if err == nil && n >= 0 && n < len(m.choices) {
				m.cursor = n
			}
		}
	}

	return m, cmd
}

func (m *model) View() string {
	b := strings.Builder{}

	b.WriteString(m.textInput.View())
	b.WriteString("\n\n")

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		digitLength := len(strconv.Itoa(len(m.choices) - 1))
		template := "%s %" + strconv.Itoa(digitLength) + "d. %s\n"
		b.WriteString(fmt.Sprintf(template, cursor, i, choice.Command.Command))
	}

	return b.String()
}
