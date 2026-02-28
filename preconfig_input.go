package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type preconfigInput struct {
	nextModel tea.Model
	flake     *flake
	options   []option
	cursor    int
}
type option struct {
	selected bool
	config   preconfiguration
}

func selectedConfigs(opts []option) []preconfiguration {
	var result []preconfiguration

	for _, opt := range opts {
		if opt.selected {
			result = append(result, opt.config)
		}
	}

	return result
}
func (n preconfigInput) Init() tea.Cmd {
	return nil
}

func (m preconfigInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.flake.preconfigs = selectedConfigs(m.options)
			return m.nextModel, nil
		}
	}
	return m, nil
}

func (m preconfigInput) View() tea.View {
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

		s.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.config.name))
	}

	s.WriteString("\nPress q to quit.\n")

	return tea.NewView(s.String())
}

var _ tea.Model = initDirenvInput(nil, nil)

func initPreconfigInput(flake *flake, nextModel tea.Model) preconfigInput {
	return preconfigInput{
		flake:     flake,
		nextModel: nextModel,
		options: []option{
			{false, goConfig()},
		},
	}
}
