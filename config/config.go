package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
		Port     string `yaml:"port"`
		SSLMode  string `yaml:"sslmode"`
		Host     string `yaml:"host"`
	}
	SV struct {
		Port string `yaml:"port"`
	}
}

func InitCfg() (*Config, error) {
	config := &Config{}

	data, err := ioutil.ReadFile("config/cfg/config.yaml")
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	log.Println("Init config", config)
	return config, nil
}
