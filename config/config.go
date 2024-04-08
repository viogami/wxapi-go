package config

import (
	"os"

	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server          ServerConfig     `yaml:"server"`
	OfficialAccount offConfig.Config `yaml:"official_account"`
	Redis           cache.RedisOpts  `yaml:"redis"`
}

type ServerConfig struct {
	Env     string `yaml:"env"`
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
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
