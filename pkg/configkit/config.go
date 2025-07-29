package configkit

import (
	"log"

	"github.com/spf13/viper"
)

var CfgFile string

type Config struct {
	Addr        string `mapstructure:"addr" default:":8080"`
	RedisURL    string `mapstructure:"redis_url"`
	DatabaseURL string `mapstructure:"database_url"`
	Ipdb        Ipdb   `mapstructure:"ipdb"`
}

type Ipdb struct {
	Ipv4 string `mapstructure:"ipv4"`
	Ipv6 string `mapstructure:"ipv6"`
}

var GlobalConfig Config

func Init() {
	// log.Println("Load config ", CfgFile)
	viper.SetConfigName(CfgFile)

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Configfile not found: " + CfgFile)
		}
	}

	err := viper.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	log.Println("Using config file:", viper.ConfigFileUsed())
	log.Println("Read config from file: ", GlobalConfig.DatabaseURL)
	log.Println("Read config from file: ", GlobalConfig.Ipdb.Ipv4)
}
