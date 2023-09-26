package config

import (
	"github.com/spf13/viper"
)

//type Email struct {
//	Name              string
//	FromEmailAddr     string
//	FromEmailPassword string
//}

type Config struct {
	//Email               Email
	EmailSenderName     string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func LoadConfig(path string) (conf Config, err error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	cfg := &Config{}
	err = viper.Unmarshal(&cfg)

	return
}
