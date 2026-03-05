package registry

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// LoadServerConfig loads the servers configuration from ~/.config/flaky-maky/servers.yaml
func LoadServerConfig() (*ServersConfig, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get config directory: %w", err)
	}

	configPath := filepath.Join(configDir, "flaky-maky", "servers.yaml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read servers config at %s: %w", configPath, err)
	}

	var config ServersConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse servers config: %w", err)
	}

	return &config, nil
}

// GetEnabledServers returns only the enabled servers
func (sc *ServersConfig) GetEnabledServers() []ServerConfig {
	var enabled []ServerConfig
	for _, server := range sc.Servers {
		if server.Enabled {
			enabled = append(enabled, server)
		}
	}
	return enabled
}
