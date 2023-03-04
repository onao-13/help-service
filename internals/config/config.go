package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Port string
}

var log = logrus.New()

func UploadDevConfig() Config {
	return loadConfig("dev")
}

func UploadProdConfig() Config {
	return loadConfig("prod")
}

func loadConfig(configName string) Config {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName(configName)
	v.SetConfigType("env")

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln("Error reading config file. Error: ", err)
	}

	var config Config

	err = v.Unmarshal(&config)
	if err != nil {
		log.Fatalln("Config unmarshal. Error: ", err)
	}

	return config
}
