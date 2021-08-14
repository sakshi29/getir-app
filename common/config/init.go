package config

import (
	"io/ioutil"
	"strings"

	"gopkg.in/gcfg.v1"
)

// Initialise application configuration
func Init(environment string) (*Config, error) {

	cfg := &Config{}

	err := cfg.ReadConfig(configPath[environment])
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// ReadConfig is file handler for reading configuration files into variable
func (cfg *Config) ReadConfig(fileName string) error {

	var configString []string

	config, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	configString = append(configString, string(config))

	err = gcfg.ReadStringInto(cfg, strings.Join(configString, "\n\n"))
	if err != nil {
		return err
	}

	return nil
}
