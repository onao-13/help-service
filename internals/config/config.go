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
	v := viper.New()
	v.SetDefault("PORT", "8082")

	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalln("Config unmarshal. Error: ", err)
	}

	return c
}

func UploadProdConfig() Config {
	v := viper.New()
	v.SetDefault("PORT", "8082")

	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalln("Config unmarshal. Error: ", err)
	}

	return c
}
