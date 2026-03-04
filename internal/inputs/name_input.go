package inputs

import (
	"fmt"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/SvBrunner/flaky-maky/internal/models"
)

type NameInput struct {
	textInput textinput.Model
	nextModel Input
	flake     *models.Flake
}

func (n *NameInput) InitInput(flake *models.Flake, nextInput Input) {
	i := textinput.New()
	i.Prompt = ""
	s := i.Styles()
	s.Cursor.Color = lipgloss.Color("63")
	i.SetStyles(s)

	i.SetWidth(48)
	i.SetValue("")
	i.CursorEnd()
	i.Focus()
	n.flake = flake
	n.textInput = i
	n.nextModel = nextInput
}

func (n NameInput) Init() tea.Cmd {
	return nil
}

func (m NameInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "enter":
			m.flake.Name = m.textInput.Value()
			return m.nextModel, nil
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m NameInput) View() tea.View {
	return tea.NewView(fmt.Sprintf(
		"\nYour name: %s\n\nPress ctrl+c to quit.",
		m.textInput.View(),
	))
}

var _ Input = (*NameInput)(nil)
