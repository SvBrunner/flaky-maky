package fileops

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SvBrunner/flaky-maky/internal/models"
	"github.com/SvBrunner/flaky-maky/internal/registry"
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
	// Just ensure the directory exists
	return os.MkdirAll(dir, 0755)
}

// SyncPreconfigs fetches configurations from configured registry servers
func SyncPreconfigs() error {
	// Load server configuration
	serversConfig, err := registry.LoadServerConfig()
	if err != nil {
		return fmt.Errorf("failed to load server configuration: %w", err)
	}

	enabledServers := serversConfig.GetEnabledServers()
	if len(enabledServers) == 0 {
		return fmt.Errorf("no enabled servers found in configuration")
	}

	configDir, err := resolveConfigPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	downloaded := 0
	upToDate := 0
	failed := 0

	for _, server := range enabledServers {
		client := registry.NewClient(server.URL, 30*time.Second)

		// Fetch list of available configs
		configs, err := client.ListConfigs()
		if err != nil {
			fmt.Printf("⚠ Warning: Failed to fetch config list from %s: %v\n", server.Name, err)
			failed++
			continue
		}

		// Download each config
		for _, config := range configs {
			filename := config.Name
			filePath := filepath.Join(configDir, filename)

			// Check if file exists locally
			var localData []byte
			if _, err := os.Stat(filePath); err == nil {
				localData, _ = os.ReadFile(filePath)
			}

			// Fetch from server (will overwrite on re-sync)
			data, _, wasModified, err := client.GetConfig(filename, localData)
			if err != nil {
				fmt.Printf("⚠ Warning: Failed to fetch %s from %s: %v\n", filename, server.Name, err)
				failed++
				continue
			}

			if !wasModified {
				upToDate++
				continue
			}

			// Write config file
			if err := os.WriteFile(filePath, data, 0644); err != nil {
				fmt.Printf("⚠ Warning: Failed to write %s: %v\n", filename, err)
				failed++
				continue
			}

			downloaded++
		}
	}

	fmt.Printf("Sync complete: %d downloaded, %d up-to-date, %d failed\n", downloaded, upToDate, failed)
	return nil
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
