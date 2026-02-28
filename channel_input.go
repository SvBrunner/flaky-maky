package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type channelInput struct {
	nextModel       tea.Model
	flake           *flake
	channels        []string
	selectedChannel int
	cursor          int
}

func (n channelInput) Init() tea.Cmd {
	return nil
}

func (m channelInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.channels)-1 {
				m.cursor++
			}

		case "space":
			m.selectedChannel = m.cursor
		case "enter":
			m.flake.channel = m.channels[m.selectedChannel]
			return m.nextModel, nil
		}
	}
	return m, nil
}

func (m channelInput) View() tea.View {
	var s strings.Builder
	s.WriteString("Which channel do you want?\n\n")

	for i, channel := range m.channels {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.selectedChannel == i {
			checked = "x"
		}

		s.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, channel))
	}

	s.WriteString("\nPress q to quit.\n")

	return tea.NewView(s.String())
}

var _ tea.Model = initDirenvInput(nil, nil)

func initChannelInput(flake *flake, nextModel tea.Model) channelInput {
	return channelInput{
		flake:     flake,
		nextModel: nextModel,
		channels: []string{
			"unstable",
			"25.11",
		},
	}
}
