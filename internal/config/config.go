package config

import (
	"log"
	"os"

	"github.com/go-playground/validator"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   `yaml:"server"`
	Database `yaml:"database"`
}

type Server struct {
	Listen string `yaml:"listen" validate:"required"`
}

type Database struct {
	User     string `yaml:"user" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
	Name     string `yaml:"name" validate:"required"`
}

func New(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Println("config file close error: ", err)
		}
	}(file)

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	if err = validator.New().Struct(config); err != nil {
		return nil, err
	}

	return config, nil
}
