package fileops

import (
	"os"
	"text/template"

	"github.com/SvBrunner/flaky-maky/internal/models"
	"github.com/SvBrunner/flaky-maky/internal/templates"
)

func GenerateFlake(flake models.Flake, path string) error {
	data := flake.ToDataModel()

	tpl, err := template.ParseFS(templates.FlakeTemplate, "flake-template.nix.tpl")

	if err != nil {
		return err
	}

	myFile, err := os.Create(path)
	if err != nil {
		return err
	}

	err = tpl.Execute(myFile, data)
	if err != nil {
		return err
	}
	err = myFile.Close()
	if err != nil {
		return err
	}

	if flake.DirenvActive {
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
