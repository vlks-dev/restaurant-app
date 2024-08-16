package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	Server struct {
		Level   string `yaml:"level"`
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Timeout struct {
			Server time.Duration `yaml:"server"`
			Write  time.Duration `yaml:"write"`
			Read   time.Duration `yaml:"read"`
			Idle   time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`
	Database struct {
		PostgreSQL struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
		} `yaml:"postgresql"`
		MongoDB struct {
			Host     string        `yaml:"host"`
			Port     string        `yaml:"port"`
			Username string        `yaml:"username"`
			Password string        `yaml:"password"`
			Database string        `yaml:"database"`
			MaxPool  int           `yaml:"max_pool"`
			MaxIdle  time.Duration `yaml:"max_idle"`
		}
	} `yaml:"database"`
}

func NewConfig(configPath string) *Config {
	config := &Config{}
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Printf("Error opening config file: %s\n", err)
		return nil
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(config); err != nil {
		fmt.Printf("Error decoding config file: %s\n", err)
		return nil
	}
	return config
}
