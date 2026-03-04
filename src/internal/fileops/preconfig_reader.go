package fileops

import (
	"os"
	"path"

	"github.com/SvBrunner/flaky-maky/internal/models"
	"gopkg.in/yaml.v3"
)

func ReadPreconfigurations(dirPath string) ([]models.Preconfiguration, error) {
	files, err := os.ReadDir(dirPath)

	if err != nil {
		return nil, err
	}

	preconfigs := make([]models.Preconfiguration, len(files))

	for i, entry := range files {

		if !entry.Type().IsRegular() {
			continue
		}
		file, err := os.ReadFile(path.Join(dirPath, entry.Name()))
		if err != nil {
			return nil, err
		}
		data := models.Preconfiguration{}

		err = yaml.Unmarshal(file, &data)
		if err := yaml.Unmarshal(file, &data); err != nil {
			return nil, err
		}
		preconfigs[i] = data
	}

	return preconfigs, nil
}
