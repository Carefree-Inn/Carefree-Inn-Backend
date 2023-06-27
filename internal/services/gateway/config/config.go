package config

import (
	"gateway/pkg/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type config struct {
	name string
}

type Config struct {
	Micro struct {
		Service string `yaml:"service"`
		Version string `yaml:"version"`
	} `yaml:"micro"`
	Gin struct {
		Mode string `yaml:"mode"`
		Port string `yaml:"port"`
	}
	QiNiu struct {
		AccessKey string `yaml:"accessKey""`
		SecretKey string `yaml:"secretKey"`
		Bucket    string `yaml:"bucket"`
		Prefix    string `json:"prefix"`
	}
}

func (cfg *config) init() {
	if cfg.name != "" {
		viper.SetConfigFile(cfg.name)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(nil, err, "读取配置文件失败:")
	}
}

func (cfg *config) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Info(nil, "Config file changed: %s", in.Name)
	})
}

func Run(name string) Config {
	cfg := config{name}
	cfg.init()
	cfg.watch()
	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal(nil, err, "配置绑定失败:")
	}
	return conf
}
