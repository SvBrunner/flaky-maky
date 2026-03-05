package registry

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// GetConfigPath returns the path to the servers configuration file
func GetConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("failed to get config directory: %w", err)
	}

	configPath := filepath.Join(configDir, "flaky-maky", "servers.yaml")
	return configPath, nil
}

// SaveServerConfig saves the servers configuration to the config file
func SaveServerConfig(config *ServersConfig) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	// Ensure the directory exists
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// AddServer adds a new server to the configuration
func AddServer(name, url string) error {
	// Load existing config or create a new one
	config, err := LoadServerConfig()
	if err != nil {
		// If config doesn't exist, create a new one
		config = &ServersConfig{Servers: []ServerConfig{}}
	}

	// Check if server with this name already exists
	for _, server := range config.Servers {
		if server.Name == name {
			return fmt.Errorf("server with name '%s' already exists", name)
		}
	}

	// Add the new server (enabled by default)
	config.Servers = append(config.Servers, ServerConfig{
		Name:    name,
		URL:     url,
		Enabled: true,
	})

	// Save the updated config
	if err := SaveServerConfig(config); err != nil {
		return err
	}

	return nil
}

// DisableServer disables a server by name
func DisableServer(name string) error {
	config, err := LoadServerConfig()
	if err != nil {
		return fmt.Errorf("failed to load server config: %w", err)
	}

	found := false
	for i, server := range config.Servers {
		if server.Name == name {
			config.Servers[i].Enabled = false
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("server with name '%s' not found", name)
	}

	// Save the updated config
	if err := SaveServerConfig(config); err != nil {
		return err
	}

	return nil
}

// EnableServer enables a server by name
func EnableServer(name string) error {
	config, err := LoadServerConfig()
	if err != nil {
		return fmt.Errorf("failed to load server config: %w", err)
	}

	found := false
	for i, server := range config.Servers {
		if server.Name == name {
			config.Servers[i].Enabled = true
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("server with name '%s' not found", name)
	}

	// Save the updated config
	if err := SaveServerConfig(config); err != nil {
		return err
	}

	return nil
}

// EnableServer enables a server by name
func DeleteServer(name string) error {
	config, err := LoadServerConfig()
	if err != nil {
		return fmt.Errorf("failed to load server config: %w", err)
	}

	found := false
	for i, server := range config.Servers {
		if server.Name == name {
			config.Servers = append(config.Servers[:i], config.Servers[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("server with name '%s' not found", name)
	}

	// Save the updated config
	if err := SaveServerConfig(config); err != nil {
		return err
	}

	return nil
}

// ListServers returns all servers in the configuration
func ListServers() ([]ServerConfig, error) {
	config, err := LoadServerConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load server config: %w", err)
	}

	return config.Servers, nil
}
