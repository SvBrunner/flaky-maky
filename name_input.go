package main

import (
	"fmt"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type nameInput struct {
	name      string
	textInput textinput.Model
	nextMdel  tea.Model
	flake     *flake
}

func (n nameInput) Init() tea.Cmd {
	return nil
}

func (m nameInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "enter":
			m.flake.name = m.textInput.Value()
			return m.nextMdel, nil
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m nameInput) View() tea.View {
	return tea.NewView(fmt.Sprintf(
		"\nYour name: %s\n\nPress ctrl+c to quit.",
		m.textInput.View(),
	))
}

var _ tea.Model = initNameInput(nil, nil)

func initNameInput(flake *flake, nextModel tea.Model) nameInput {
	i := textinput.New()
	i.Prompt = ""
	s := i.Styles()
	s.Cursor.Color = lipgloss.Color("63")
	i.SetStyles(s)

	i.SetWidth(48)
	i.SetValue("")
	i.CursorEnd()
	i.Focus()

	return nameInput{
		textInput: i,
		nextMdel:  nextModel,
		flake:     flake,
	}
}
