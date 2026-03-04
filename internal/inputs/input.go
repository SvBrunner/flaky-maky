package inputs

import (
	tea "charm.land/bubbletea/v2"
	"github.com/SvBrunner/flaky-maky/internal/models"
)

type Input interface {
	tea.Model
	InitInput(flake *models.Flake, nextInput Input)
}
