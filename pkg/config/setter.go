package config

import (
	"github.com/spf13/viper"
	"log"
)

func Set() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

}
