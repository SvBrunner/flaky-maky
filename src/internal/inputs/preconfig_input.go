package inputs

import (
	"fmt"
	"log"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/SvBrunner/flaky-maky/internal/fileops"
	"github.com/SvBrunner/flaky-maky/internal/models"
)

type PreConfigInput struct {
	nextModel tea.Model
	flake     *models.Flake
	options   []option
	cursor    int
}

func (n *PreConfigInput) InitInput(flake *models.Flake, nextInput Input) {
	n.flake = flake
	n.nextModel = nextInput
	preconfigs, err := fileops.ReadPreconfigurations()
	if err != nil {
		log.Fatal(err)
	}
	n.options = make([]option, len(preconfigs))
	for i, cfg := range preconfigs {
		n.options[i] = option{false, cfg}
	}
}

type option struct {
	selected bool
	config   models.Preconfiguration
}

func selectedConfigs(opts []option) []models.Preconfiguration {
	var result []models.Preconfiguration

	for _, opt := range opts {
		if opt.selected {
			result = append(result, opt.config)
		}
	}

	return result
}
func (n PreConfigInput) Init() tea.Cmd {
	return nil
}

func (m PreConfigInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.flake.Preconfigs = selectedConfigs(m.options)
			return m.nextModel, nil
		}
	}
	return m, nil
}

func (m PreConfigInput) View() tea.View {
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

		s.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.config.Name))
	}

	s.WriteString("\nPress q to quit.\n")

	return tea.NewView(s.String())
}

var _ Input = (*PreConfigInput)(nil)
