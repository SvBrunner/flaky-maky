package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type systemsInput struct {
	nextModel tea.Model
	flake     *flake
	options   []systemOption
	cursor    int
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
func (n systemsInput) Init() tea.Cmd {
	return nil
}

func (m systemsInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.flake.systems = selectedSystems(m.options)
			return m.nextModel, nil
		}
	}
	return m, nil
}

func (m systemsInput) View() tea.View {
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

var _ tea.Model = initDirenvInput(nil, nil)

func initSystemsInput(flake *flake, nextModel tea.Model) systemsInput {
	return systemsInput{
		flake:     flake,
		nextModel: nextModel,
		options: []systemOption{
			{false, "x86_64-linux"},
			{false, "aarch64-linux"},
			{false, "x86_64-darwin"},
			{false, "aarch64-darwin"},
		},
	}
}
