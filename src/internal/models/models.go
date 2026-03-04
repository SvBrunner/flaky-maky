package models

import "fmt"

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

func (flake Flake) ToString() string {
	s := fmt.Sprintf("Flake %s\n", flake.Name)
	s += fmt.Sprintf("Channel : %s\n", flake.Channel)
	s += fmt.Sprintf("Direnv active : %t\n", flake.DirenvActive)
	s += "Preconfigurations : \n"
	for _, config := range flake.Preconfigs {
		s += fmt.Sprintf("%s\n", config.Name)
	}
	return s
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
	Name        string
	Packages    []string
	Environment []EnvVariable
	Shellhook   string
}

type EnvVariable struct {
	Name  string
	Value string
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
