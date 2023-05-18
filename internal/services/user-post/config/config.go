package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type config struct {
	name string
}

type Config struct {
	Database struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"database"`
	Server struct {
		Http struct {
			Address string `yaml:"address"`
		} `yaml:"http"`
	} `yaml:"server"`
	Micro struct {
		Service string `yaml:"service"`
		Version string `yaml:"version"`
	} `yaml:"micro"`
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
		log.Println(nil, "Config file changed: %s", in.Name)
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
