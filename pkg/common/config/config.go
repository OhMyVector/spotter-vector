package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

type Configuration struct {
	Server *Server      `yaml:"server,omitempty"`
	App    *Application `yaml:"application,omitempty"`
}

type Server struct {
	Port         string `yaml:"port,omitempty"`
	Debug        bool   `yaml:"debug,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
}

type Application struct {
	StaticPath string `yaml:"static_path,omitempty"`
}
