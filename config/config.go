package config

import (
	"github.com/spf13/viper"
)

//TODO verificar se vai ser usado
//Load is a func that returns the configuration fetched from the env file.
//func Load() (*viper.Viper, error) {
//	conf := viper.GetViper()
//	conf.AddConfigPath(".")
//	conf.SetConfigFile(".env")
//	conf.SetConfigType("env")
//	err := conf.ReadInConfig()
//	if err != nil {
//		return nil, err
//	}
//	return conf, nil
//}

//Load is a function that returns the CSV header configuration
func LoadConfigHeader() (*viper.Viper, error) {
	conf := viper.GetViper()
	conf.AddConfigPath(".")
	conf.SetConfigFile("headerConfiguration")
	conf.SetConfigType("json")
	err := conf.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return conf, nil
}
