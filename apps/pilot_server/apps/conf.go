package apps

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type AgentConfig struct {
	Addr string `yaml:"addr"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DBConfig struct {
	Addr   string `yaml:"addr"`
	Port   string `yaml:"port"`
	User   string `yaml:"user"`
	PW     string `yaml:"passwd"`
	DBName string `yaml:"db"`
}

type Config struct {
	Version string       `yaml:"version"`
	Agent   AgentConfig  `yaml:"pilot_agent"`
	Server  ServerConfig `yaml:"pilot_server"`
	DB      DBConfig     `yaml:"pilot_db"`
}

var (
	Conf Config
)

func ReadConfig() error {
	rfile, err := ioutil.ReadFile("./conf/config.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(rfile, &Conf)
	if err != nil {
		return err
	}

	fmt.Println("#1")

	return nil
}
