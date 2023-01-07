package apps

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type AgentConfig struct {
	Port    string `yaml:"port"`
	Version string `yaml:"version"`
}

type ServerConfig struct {
	Addr string `yaml:"addr"`
}

type Config struct {
	Agent  AgentConfig  `yaml:"pilot_agent"`
	Server ServerConfig `yaml:"pilot_server"`
}

var (
	Conf Config
)

func ReadConfig() error {
	rfile, err := ioutil.ReadFile("./conf/config.yaml")
	if err != nil {
		return err
	}

	var datas map[string]interface{}
	err = yaml.Unmarshal(rfile, &datas)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(rfile, &Conf)
	if err != nil {
		return err
	}

	if 0 == len(Conf.Agent.Port) {
		return errors.New("Agent post is empty")
	}

	if 0 == len(Conf.Server.Addr) {
		return errors.New("Server addr is empty")
	}

	Logs.Info(fmt.Sprintf("pilot_agent version %s", Conf.Agent.Version))

	return nil
}
