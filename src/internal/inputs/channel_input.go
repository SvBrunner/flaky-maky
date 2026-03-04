package inputs

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/SvBrunner/flaky-maky/internal/models"
)

type channel struct {
	name string
	url  string
}

type ChannelInput struct {
	nextModel       Input
	flake           *models.Flake
	channels        []channel
	selectedChannel int
	cursor          int
}

func (n *ChannelInput) InitInput(flake *models.Flake, nextInput Input) {
	n.flake = flake
	n.nextModel = nextInput
	n.channels = []channel{
		{name: "unstable", url: "github:NixOS/nixpkgs/nixos-unstable"},
		{name: "25.11", url: "github:NixOS/nixpkgs/nixos-25.11"},
	}

}

func (n ChannelInput) Init() tea.Cmd {
	return nil
}

func (m ChannelInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.flake.Channel = m.channels[m.selectedChannel].url
			return m.nextModel, nil
		}
	}
	return m, nil
}

func (m ChannelInput) View() tea.View {
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

		s.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, channel.name))
	}

	s.WriteString("\nPress q to quit.\n")

	return tea.NewView(s.String())
}

var _ Input = (*ChannelInput)(nil)
