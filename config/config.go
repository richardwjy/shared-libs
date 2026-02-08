package config

import (
	"fmt"
	"os"
	"path/filepath"

	"go.yaml.in/yaml/v4"
)

type Config interface {
	Validate() error
}

func Load(path string, config Config) error {
	// Clean path to avoid traversal weirdness
	cleanPath := filepath.Clean(path)

	info, err := os.Stat(cleanPath)
	if err != nil {
		return fmt.Errorf("stat config file: %w", err)
	}

	if info.IsDir() {
		return fmt.Errorf("config path is a directory: %s", cleanPath)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return fmt.Errorf("unmarshal yaml: %w", err)
	}

	return config.Validate()
}
