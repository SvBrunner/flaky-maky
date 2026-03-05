package registry

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client handles HTTP requests to a registry server
type Client struct {
	BaseURL string
	Timeout time.Duration
}

// NewClient creates a new registry client
func NewClient(baseURL string, timeout time.Duration) *Client {
	return &Client{
		BaseURL: baseURL,
		Timeout: timeout,
	}
}

// ListConfigs fetches the list of available configs from the server
func (c *Client) ListConfigs() ([]ConfigMetadata, error) {
	url := c.BaseURL + "/configs"

	client := &http.Client{Timeout: c.Timeout}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch config list: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		var errResp ConfigError
		if err := json.Unmarshal(body, &errResp); err == nil {
			return nil, fmt.Errorf("server error: %s - %s", errResp.Code, errResp.Message)
		}
		return nil, fmt.Errorf("server returned status %d", resp.StatusCode)
	}

	var configs []ConfigMetadata
	if err := json.NewDecoder(resp.Body).Decode(&configs); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return configs, nil
}

// GetConfig fetches a specific config file from the server
// localData is the bytes of an existing local file (can be nil)
// Returns (data, etag, wasModified, error)
func (c *Client) GetConfig(filename string, localData []byte) ([]byte, string, bool, error) {
	url := c.BaseURL + "/configs/" + filename

	client := &http.Client{Timeout: c.Timeout}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", false, fmt.Errorf("failed to create request: %w", err)
	}

	// If local file exists, send SHA-256 hash as If-None-Match
	if len(localData) > 0 {
		hash := sha256.Sum256(localData)
		etag := fmt.Sprintf("\"%s\"", "sha256:"+hex.EncodeToString(hash[:]))
		req.Header.Set("If-None-Match", etag)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, "", false, fmt.Errorf("failed to fetch config: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	// 304 Not Modified
	if resp.StatusCode == http.StatusNotModified {
		return localData, "", false, nil
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		var errResp ConfigError
		if err := json.Unmarshal(body, &errResp); err == nil {
			return nil, "", false, fmt.Errorf("server error: %s - %s", errResp.Code, errResp.Message)
		}
		return nil, "", false, fmt.Errorf("server returned status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", false, fmt.Errorf("failed to read response: %w", err)
	}

	etag := resp.Header.Get("ETag")
	return data, etag, true, nil
}
