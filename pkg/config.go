package pkg

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// URLConfig holds the structure of each URL configuration.
type URLConfig struct {
	Name    string `yaml:"name"`
	Timeout int    `yaml:"timeout"`
	Crawl   bool   `yaml:"crawl"`
	Enabled bool   `yaml:"enabled"`
}

// Config represents the top-level configuration structure.
type Config struct {
	URLs []URLConfig `yaml:"urls"`
}

func NewConfig(configPath string) Config {
	var config Config
	// #nosec
	yamlContent, err := os.ReadFile(configPath)

	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal([]byte(yamlContent), &config)
	if err != nil {
		fmt.Println("Error parsing YAML file:", err)
		os.Exit(1)
	}
	return config
}
