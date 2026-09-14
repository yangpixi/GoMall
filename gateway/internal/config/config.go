package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Server Server `yaml:"server"`
	JWT    JWT    `yaml:"jwt"`
}

type Server struct {
	Port int `yaml:"port"`
}

type JWT struct {
	SecretKey string `yaml:"secret-key"`
}

func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	c := &Config{}

	if err = yaml.Unmarshal(b, c); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if c.JWT.SecretKey == "" {
		return nil, errors.New("missing secretKey")
	}

	return c, nil
}
