package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

// Config holds all the user options for editor, folder locations, etc...
type Config struct {
	Editor        string `yaml:"editor"`
	PublicFolder  string `yaml:"public_folder"`
	PrivateFolder string `yaml:"private_folder"`
}

// GetConfig return an struct reading from the config yaml file
func (c *Config) GetConfig() (*Config, error) {
	configPath := c.getConfigFilePath()
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) getConfigFilePath() string {
	base, _ := os.UserHomeDir()
	return path.Join(base, ".config", "remarke", "config.yaml")
}

func (c *Config) getConfigFolderPath() string {
	base, _ := os.UserHomeDir()
	return path.Join(base, ".config", "remarke")
}

func (c *Config) setConfigFile(data []byte) (bool, error) {
	configPath := c.getConfigFolderPath()
	_, err := ioutil.ReadFile(configPath)

	if err != nil {
		os.Mkdir(configPath, 0755)
		err = ioutil.WriteFile(fmt.Sprintf("%s/config.yaml", configPath), data, 0644)

		if err != nil {
			return false, err
		}
	}

	return true, nil
}
