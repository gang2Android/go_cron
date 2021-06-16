package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Name  string `yaml:"name"`
	Port  string `yaml:"port"`
	Mysql Mysql  `yaml:"mysql"`
	Redis Redis  `yaml:"redis"`
	Logs  Logs   `yaml:"logs"`
}

type Mysql struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Db   string `yaml:"db"`
	Name string `yaml:"name"`
	Pwd  string `yaml:"pwd"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Pwd  string `yaml:"pwd"`
	Db   int    `yaml:"db"`
}

type Logs struct {
	FilePath   string `yaml:"file_path"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
	Debug      bool   `yaml:"debug"`
}

func LoadConfig(configPath string) *Config {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	config := &Config{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}
	return config
}
