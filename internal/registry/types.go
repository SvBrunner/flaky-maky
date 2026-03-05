package registry

import "time"

// ConfigMetadata represents metadata about a config file from the server
type ConfigMetadata struct {
	Name         string    `json:"name"`
	LastModified time.Time `json:"lastModified"`
	ETag         string    `json:"etag"`
}

// ConfigError represents an error response from the server
type ConfigError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ServerConfig represents a single registry server configuration
type ServerConfig struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	Enabled bool   `yaml:"enabled"`
}

// ServersConfig represents the full servers configuration file
type ServersConfig struct {
	Servers []ServerConfig `yaml:"servers"`
}
