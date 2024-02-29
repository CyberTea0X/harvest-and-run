package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type ClientConfig struct {
	ServerIp   string `toml:"server_ip"`
	ServerPort uint   `toml:"server_port"`
}

func (c *ClientConfig) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.ServerIp, c.ServerPort)
}

type HostConfig struct {
	Port uint
}

func (c *HostConfig) Address() string {
	return fmt.Sprintf("localhost:%d", c.Port)
}

type Config struct {
	Client ClientConfig
	Host   HostConfig
}

func FromToml(path string) (*Config, error) {
	c := new(Config)
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Join(errors.New("Failed to read config"), err)
	}
	err = toml.Unmarshal(file, c)
	if err != nil {
		return nil, errors.Join(errors.New("Failed to unmarshal config"), err)
	}

	return c, nil
}
