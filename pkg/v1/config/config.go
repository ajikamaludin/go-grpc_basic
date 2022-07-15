package config

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env    string `yaml:"env"`
	Server struct {
		Host string `yaml:"host"`
		Grpc struct {
			Port int `yaml:"port"`
		} `yaml:"grpc"`
		Rest struct {
			Port int `yaml:"port"`
		} `yaml:"rest"`
	} `yaml:"server"`
}

func New() (*Config, error) {
	cfgPath, err := parseFlag()
	if err != nil {
		log.Fatal(err)
	}

	config := &Config{}

	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	err = validateConfigData(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func parseFlag() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config.yaml", "path to config file")

	// Actually parse the flags
	flag.Parse()

	//Validate the path first
	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func validateConfigPath(path string) error {
	s, err := os.Stat(path)

	if err != nil {
		return err
	}

	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a config file", path)
	}

	return nil
}

func validateConfigData(config *Config) error {
	if config.Env == "" {
		return errors.New("env is empty")
	}
	if config.Server.Host == "" {
		return errors.New("server.host is empty")
	}
	if config.Server.Grpc.Port == 0 {
		return errors.New("server.grpc.port is empty")
	}
	if config.Server.Rest.Port == 0 {
		return errors.New("server.rest.port is empty")
	}

	return nil
}
