package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type flakeModel struct {
	flake *flake
}

func initialModel(flake *flake) flakeModel {
	return flakeModel{
		flake: flake,
	}
}

func (m flakeModel) Init() tea.Cmd {
	return nil
}

func (m flakeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m flakeModel) View() tea.View {
	s := m.flake.toString()
	s += "\nPress q to quit.\n"
	return tea.NewView(s)
}

func initInputs() tea.Model {
	flake := flake{}
	final := initialModel(&flake)
	preConfigInput := initPreconfigInput(&flake, final)
	direnvInput := initDirenvInput(&flake, preConfigInput)
	channelInput := initChannelInput(&flake, direnvInput)
	nameInput := initNameInput(&flake, channelInput)
	return nameInput

}
func main() {
	p := tea.NewProgram(initInputs())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
