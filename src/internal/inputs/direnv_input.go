package inputs

import (
	tea "charm.land/bubbletea/v2"
	"github.com/SvBrunner/flaky-maky/internal/models"
)

type DirenvInput struct {
	nextModel Input
	flake     *models.Flake
}

func (n *DirenvInput) InitInput(flake *models.Flake, nextInput Input) {
	n.flake = flake
	n.nextModel = nextInput
}

func (n DirenvInput) Init() tea.Cmd {
	return nil
}

func (m DirenvInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "y":
			m.flake.DirenvActive = true
			return m.nextModel, nil
		case "n":
			m.flake.DirenvActive = false
			return m.nextModel, nil
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m DirenvInput) View() tea.View {
	return tea.NewView(
		"\nDo you want a direnv in this directory? (y/n)\n\nPress q to quit.",
	)
}

var _ Input = (*DirenvInput)(nil)
