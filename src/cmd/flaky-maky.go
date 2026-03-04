package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/SvBrunner/flaky-maky/internal/fileops"
	"github.com/SvBrunner/flaky-maky/internal/inputs"
	"github.com/SvBrunner/flaky-maky/internal/models"
)

func initInputs() tea.Model {
	flake := models.Flake{}
	inputs := []inputs.Input{
		&inputs.NameInput{},
		&inputs.ChannelInput{},
		&inputs.SystemInput{},
		&inputs.PreConfigInput{},
		&inputs.DirenvInput{},
		&inputs.FinalInput{},
	}
	for index, input := range inputs {
		if len(inputs) > index+1 {
			input.InitInput(&flake, inputs[index+1])
		} else {
			input.InitInput(&flake, nil)
		}
	}

	return inputs[0]

}
func main() {
	fileops.PopulatePreconfigs()
	p := tea.NewProgram(initInputs())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
