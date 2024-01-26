package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Wx     WxConfig     `yaml:"wx"`
}

type ServerConfig struct {
	Env     string `yaml:"env"`
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type WxConfig struct {
	WxToken   string `yaml:"wxToken"`
	AppID     string `yaml:"appID"`
	AppSecret string `yaml:"appSecret"`
}

// 读取配置文件 app.yaml
func ConfigParse(appConfig *string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(*appConfig)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
