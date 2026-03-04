package models

import (
	"fmt"
	"strings"
)

type Flake struct {
	Name         string
	Preconfigs   []Preconfiguration
	Channel      string
	DirenvActive bool
	Systems      []string
}

type FlakeData struct {
	Description          string
	Channel              string
	Packages             []string
	Systems              []string
	ShellHooks           []string
	EnvironmentVariables []EnvVariable
}

func (f Flake) ToString() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Flake: %s\n", f.Name)
	fmt.Fprintf(&b, "Channel: %s\n", f.Channel)
	fmt.Fprintf(&b, "Direnv:  %t\n", f.DirenvActive)
	b.WriteString("Systems:\n")
	for _, s := range f.Systems {
		fmt.Fprintf(&b, "  • %s\n", s)
	}
	b.WriteString("Preconfigurations:\n")
	for _, p := range f.Preconfigs {
		fmt.Fprintf(&b, "  • %s\n", p.Name)
	}

	return b.String()
}

func (flake Flake) ToDataModel() FlakeData {
	totalPackages := 0
	totalEnvVariables := 0
	for _, preconfig := range flake.Preconfigs {
		totalPackages += len(preconfig.Packages)
		totalEnvVariables += len(preconfig.Environment)
	}

	allPackages := make([]string, 0, totalPackages)
	allShellhooks := make([]string, 0, len(flake.Preconfigs))
	allEnvVariables := make([]EnvVariable, 0, totalEnvVariables)

	for _, preconfig := range flake.Preconfigs {
		allPackages = append(allPackages, preconfig.Packages...)
		allShellhooks = append(allShellhooks, preconfig.Shellhook)
		allEnvVariables = append(allEnvVariables, preconfig.Environment...)
	}
	return FlakeData{
		Channel:              flake.Channel,
		Description:          flake.Name,
		Systems:              flake.Systems,
		Packages:             allPackages,
		ShellHooks:           allShellhooks,
		EnvironmentVariables: allEnvVariables,
	}
}

type Preconfiguration struct {
	Name        string        `yaml:"Name"`
	Packages    []string      `yaml:"Packages"`
	Environment []EnvVariable `yaml:"Environment"`
	Shellhook   string        `yaml:"Shellhook"`
}

type EnvVariable struct {
	Name  string `yaml:"Name"`
	Value string `yaml:"Value"`
}

func GoConfig() Preconfiguration {
	return Preconfiguration{
		Name:     "Go dev environment",
		Packages: []string{"go", "gopls"},
		Environment: []EnvVariable{
			{
				Name:  "GOPATH",
				Value: "/home/sven/go",
			},
		},
		Shellhook: "go version",
	}
}
