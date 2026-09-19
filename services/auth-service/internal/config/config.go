package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-yaml"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	JWT      JWT      `yaml:"jwt"`
}

type Server struct {
	Port int `yaml:"port"`
}

type Database struct {
	Connection Connection `yaml:"connection"`
	User       string     `yaml:"user" validate:"required"`
	Password   string     `yaml:"password" validate:"required"`
	DB         string     `yaml:"db" validate:"required"`
}

type Connection struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type JWT struct {
	SecretKey         string `yaml:"secret-key" validate:"required"`
	Expiration        int    `yaml:"expiration" validate:"required"`
	RefreshExpiration int    `yaml:"refresh-expiration" validate:"required"`
}

// Load config for server port, database etc
func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// setting default value for service
	c := &Config{
		Server: Server{Port: 8081},
		Database: Database{
			Connection: Connection{
				Port: 5432,
				Host: "127.0.0.1",
			},
		},
	}

	if err := yaml.UnmarshalWithOptions(b, c, yaml.Validator(validator.New())); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if c.JWT.SecretKey == "" {
		return nil, errors.New("missing secretKey")
	}

	return c, nil
}
