package main

import "fmt"

type flake struct {
	name         string
	preconfigs   []preconfiguration
	channel      string
	direnvActive bool
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

type preconfiguration struct {
	name        string
	packages    []string
	environment []envVariable
	shellhook   string
}

type envVariable struct {
	name  string
	value string
}

func goConfig() preconfiguration {
	return preconfiguration{
		name:     "Go dev environment",
		packages: []string{"go", "gopls"},
		environment: []envVariable{
			{
				name:  "GOPATH",
				value: "/home/sven/go",
			},
		},
		shellhook: "go version",
	}
}
