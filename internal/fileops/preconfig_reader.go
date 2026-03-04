package fileops

import (
	"os"
	"path/filepath"

	"github.com/SvBrunner/flaky-maky/internal/models"
	"github.com/SvBrunner/flaky-maky/internal/templates"
	"gopkg.in/yaml.v3"
)

func resolveConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "flaky-maky", "templates"), nil
}

func PopulatePreconfigs() error {
	dir, err := resolveConfigPath()
	if err != nil {
		return err
	}
	os.MkdirAll(dir, 0755)
	entries, err := templates.DefaultTemplates.ReadDir("templates")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		dst := filepath.Join(dir, entry.Name())

		if _, err := os.Stat(dst); err == nil {
			continue
		}

		data, err := templates.DefaultTemplates.ReadFile("templates/" + entry.Name())
		if err != nil {
			return err
		}

		err = os.WriteFile(dst, data, 0644)
		if err != nil {
			return err
		}
	}
	return err
}

func ReadPreconfigurations() ([]models.Preconfiguration, error) {
	dirPath, err := resolveConfigPath()
	if err != nil {
		return nil, err
	}
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	preconfigs := make([]models.Preconfiguration, 0, len(entries))

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		b, err := os.ReadFile(filepath.Join(dirPath, entry.Name()))
		if err != nil {
			return nil, err
		}

		var data models.Preconfiguration
		if err := yaml.Unmarshal(b, &data); err != nil {
			return nil, err
		}

		preconfigs = append(preconfigs, data)
	}

	return preconfigs, nil
}
