package config

import (
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

//CONFIG model for reading configurations
type CONFIG struct {
	PostgresConnectionURL string `json:"_" mapstructure:"POSTGRES_CONNECTION_URL"`

	PORT string `json:"_" mapstructure:"PORT"`
}

//LoadConfig reads configurations from app.env file
func LoadConfig(path string) (*CONFIG, error) {
	viper.AddConfigPath(path)
	viper.AddConfigPath("../" + path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	var config CONFIG

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

//Config is initialized pointer to configuration settings
var Config *CONFIG

func init() {
	cfg, er := LoadConfig("./config")
	if er != nil {
		log.Fatalln(errors.Wrap(er, "unable to read .env file"))
	}

	Config = cfg
}
