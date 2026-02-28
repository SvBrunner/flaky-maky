package main

import (
	tea "charm.land/bubbletea/v2"
)

type direnvInput struct {
	nextModel tea.Model
	flake     *flake
}

func (n direnvInput) Init() tea.Cmd {
	return nil
}

func (m direnvInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "y":
			m.flake.direnvActive = true
			return m.nextModel, nil
		case "n":
			m.flake.direnvActive = false
			return m.nextModel, nil
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m direnvInput) View() tea.View {
	return tea.NewView(
		"\nDo you want a direnv in this directory? (y/n)\n\nPress ctrl+c to quit.",
	)
}

var _ tea.Model = initDirenvInput(nil, nil)

func initDirenvInput(flake *flake, nextModel tea.Model) direnvInput {
	return direnvInput{
		flake:     flake,
		nextModel: nextModel,
	}
}
