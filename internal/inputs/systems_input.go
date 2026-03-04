package inputs

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/SvBrunner/flaky-maky/internal/models"
)

type SystemInput struct {
	nextModel tea.Model
	flake     *models.Flake
	options   []systemOption
	cursor    int
}

func (n *SystemInput) InitInput(flake *models.Flake, nextInput Input) {
	n.flake = flake
	n.nextModel = nextInput
	n.options = []systemOption{
		{false, "x86_64-linux"},
		{false, "aarch64-linux"},
		{false, "x86_64-darwin"},
		{false, "aarch64-darwin"},
	}
}

type systemOption struct {
	selected bool
	system   string
}

func selectedSystems(opts []systemOption) []string {
	var result []string

	for _, opt := range opts {
		if opt.selected {
			result = append(result, opt.system)
		}
	}

	return result
}
func (n SystemInput) Init() tea.Cmd {
	return nil
}

func (m SystemInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}

		case "space":
			m.options[m.cursor].selected = !m.options[m.cursor].selected
		case "enter":
			m.flake.Systems = selectedSystems(m.options)
			return m.nextModel, nil
		}
	}
	return m, nil
}

func (m SystemInput) View() tea.View {
	var s strings.Builder
	s.WriteString("Which preconfigurations do you want?\n\n")

	for i, choice := range m.options {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.options[i].selected {
			checked = "x"
		}

		s.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.system))
	}

	s.WriteString("\nPress q to quit.\n")

	return tea.NewView(s.String())
}

var _ Input = (*SystemInput)(nil)
