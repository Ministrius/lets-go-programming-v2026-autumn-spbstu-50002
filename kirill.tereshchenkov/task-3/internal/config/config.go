package config

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)

	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("failed to decode config: %w", err)
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" {
		return Config{}, fmt.Errorf("input-file is empty")
	}

	if cfg.OutputFile == "" {
		return Config{}, fmt.Errorf("output-file is empty")
	}

	return cfg, nil
}
