package models
import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"errors"
	"fmt"
)

type AppConfig struct {
	Mongodb MongodbConfig `yaml:"mongodb"`
	Logconf string        `yaml:"log_config"`
}

type MongodbConfig struct {
	Host string   `yaml:"host"`
	Db   string   `yaml:"db"`
}

func (mc *MongodbConfig) isValid() bool {
	return len(mc.Host) >0 && len(mc.Db) >0
}

func (ac *AppConfig) isValid() bool {
	return ac.Mongodb.isValid()
}

var DBConfig AppConfig

func ParseConfigFile(filepath string) error {
	if config, err := ioutil.ReadFile(filepath); err == nil {
		if err = yaml.Unmarshal(config, &DBConfig); err != nil {
			return err
		}
		if ! DBConfig.isValid() {
			return errors.New("Invalid configuration!")
		}
	}else {
		fmt.Println("parse error ")
		return err
	}
	return nil
}