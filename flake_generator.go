package main

import (
	"os"
	"strings"
	"text/template"
)

func generateFlake(flake flake, path string) error {
	data := flake.toDataModel()
	funcMap := template.FuncMap{
		"join": strings.Join,
	}

	tpl, err := template.ParseFiles("flake-template.nix.tpl")

	if err != nil {
		return err
	}

	myFile, err := os.Create(path)
	if err != nil {
		return err
	}
	tpl.Funcs(funcMap)

	err = tpl.Execute(myFile, data)
	if err != nil {
		return err
	}
	err = myFile.Close()
	if err != nil {
		return err
	}

	if flake.direnvActive {
		err = generateDirenv()
	}
	return err

}

func generateDirenv() error {
	direnvFile, err := os.Create(".envrc")
	if err != nil {
		return err
	}
	_, err = direnvFile.WriteString("use flake")
	return direnvFile.Close()
}
