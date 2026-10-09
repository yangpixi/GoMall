package config

import (
	"bytes"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-yaml"
)

type Config struct {
	Server   Server
	Database Database
	JWT      JWT
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
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type JWT struct {
	SecretKey string `yaml:"secret-key" validate:"required"`
}

// Load the very config
func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if len(b) == 0 {
		return nil, fmt.Errorf("config is empty")
	}

	// setting default value for service
	c := &Config{
		Server: Server{Port: 8085},
		Database: Database{
			Connection: Connection{
				Port: 5432,
				Host: "127.0.0.1",
			},
		},
	}

	dec := yaml.NewDecoder(
		bytes.NewReader(b),
		yaml.Validator(validator.New()),
		yaml.Strict(),
	)

	err = dec.Decode(c)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return c, nil

}
