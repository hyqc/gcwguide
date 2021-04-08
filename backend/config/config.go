package conf

import (
	"errors"
	"gcwguide/extend/util"
	"gopkg.in/yaml.v2"
)

// Config server config
type Config struct {
	Server Server `yaml:"server"`
	Store  Store  `yaml:"store"`
}

// App application config
type Server struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

// Store store config
type Store struct {
	Drive string `yaml:"drive"`
	Type  string `yaml:"type"`
	Path  string `yaml:"path"`
}

func InitConfig() (*Config, error) {
	config, err := util.ParseYaml("config/config.yaml")
	if err != nil {
		errMsg := errors.New("parse config.yaml file error : " + err.Error())
		return nil, errMsg
	}
	var c Config
	if err := yaml.Unmarshal(config, &c); err != nil {
		errMsg := errors.New("parse config []byte to struct error ：" + err.Error())
		return nil, errMsg
	}
	return &c, nil
}
