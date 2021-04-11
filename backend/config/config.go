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
	Name    string        `yaml:"name"`
	Port    string        `yaml:"port"`
}

// Store store config
type Store struct {
	Drive      string `yaml:"drive"`
	Type       string `yaml:"type"`
	Path       string `yaml:"path"`
	Suffix     string `yaml:"suffix"`
	BackupsDir string `json:"backups_dir"`
	BackupsMax int    `json:"backups_max"`
	FileSync   *util.FileSync
}

var App Config

const (
	StoreDriveFile = "file"
)

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
	if err := c.initStoreDrive(); err != nil {
		return &c, nil
	}
	return &c, nil
}

// initStoreDrive 初始化存储驱动
func (c *Config) initStoreDrive() (err error) {
	switch c.Store.Drive {
	case StoreDriveFile:
		fallthrough
	default:
		fileSync := util.FileSync{}
		fileSync.FilePath = c.Store.Path
		if err = fileSync.InitStoreFile(c.Store.Path, 0755); err != nil {
			return err
		}
		c.Store.FileSync = &fileSync
		if c.Store.BackupsMax <= 7 {
			c.Store.BackupsMax = 7
		}
	}
	return nil
}
