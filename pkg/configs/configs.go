package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var Cfg *Configs

type Configs struct {
	JWT_SECRET_KEY string
	MONGODB_URI    string
}

func LoadConfigs() (*Configs, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	config := &Configs{
		JWT_SECRET_KEY: viper.GetString("JWT_SECRET_KEY"),
		MONGODB_URI:    viper.GetString("MONGODB_URI"),
	}

	Cfg = config

	return config, nil
}
