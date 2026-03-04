package inputs

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/SvBrunner/flaky-maky/internal/fileops"
	"github.com/SvBrunner/flaky-maky/internal/models"
)

type FinalInput struct {
	flake *models.Flake
	err   string
}

func (f FinalInput) Init() tea.Cmd {
	return nil
}

func (m FinalInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			err := fileops.GenerateFlake(*m.flake, "flake.nix")
			if err != nil {
				m.err = err.Error()
			} else {
				return m, tea.Quit
			}
		}

	}
	return m, nil
}

func (m FinalInput) View() tea.View {
	s := m.flake.ToString()
	if m.err != "" {
		s += fmt.Sprintf("Error: %s\n", m.err)
	}
	s += "\nPress enter to create the flake\n"
	s += "\nPress q to quit.\n"
	return tea.NewView(s)
}

func (f *FinalInput) InitInput(flake *models.Flake, nextInput Input) {
	f.flake = flake
}

var _ Input = (*FinalInput)(nil)
