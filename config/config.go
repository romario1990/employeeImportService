package config

import (
	"github.com/spf13/viper"
)

//Load is a function that returns the csv header configuration
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
