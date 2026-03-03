package main

import "fmt"

type flake struct {
	name         string
	preconfigs   []preconfiguration
	channel      string
	direnvActive bool
	systems      []string
}

type flakeData struct {
	Description          string
	Channel              string
	Packages             []string
	Systems              []string
	ShellHooks           []string
	EnvironmentVariables []envVariable
}

func (flake flake) toString() string {
	s := fmt.Sprintf("Flake %s\n", flake.name)
	s += fmt.Sprintf("Channel : %s\n", flake.channel)
	s += fmt.Sprintf("Direnv active : %t\n", flake.direnvActive)
	s += "Preconfigurations : \n"
	for _, config := range flake.preconfigs {
		s += fmt.Sprintf("%s\n", config.name)
	}
	return s
}

func (flake flake) toDataModel() flakeData {
	totalPackages := 0
	totalEnvVariables := 0
	for _, preconfig := range flake.preconfigs {
		totalPackages += len(preconfig.packages)
		totalEnvVariables += len(preconfig.environment)
	}

	allPackages := make([]string, 0, totalPackages)
	allShellhooks := make([]string, 0, len(flake.preconfigs))
	allEnvVariables := make([]envVariable, 0, totalEnvVariables)

	for _, preconfig := range flake.preconfigs {
		allPackages = append(allPackages, preconfig.packages...)
		allShellhooks = append(allShellhooks, preconfig.shellhook)
		allEnvVariables = append(allEnvVariables, preconfig.environment...)
	}
	return flakeData{
		Channel:              flake.channel,
		Description:          flake.name,
		Systems:              flake.systems,
		Packages:             allPackages,
		ShellHooks:           allShellhooks,
		EnvironmentVariables: allEnvVariables,
	}
}

type preconfiguration struct {
	name        string
	packages    []string
	environment []envVariable
	shellhook   string
}

type envVariable struct {
	Name  string
	Value string
}

func goConfig() preconfiguration {
	return preconfiguration{
		name:     "Go dev environment",
		packages: []string{"go", "gopls"},
		environment: []envVariable{
			{
				Name:  "GOPATH",
				Value: "/home/sven/go",
			},
		},
		shellhook: "go version",
	}
}
