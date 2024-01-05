package config

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)


type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port int
	RunMode string
}

func GetConfig() *Config {
	configPath := GetConfigPath(os.Getenv("APP_ENV"))

	configViper, err := LoadConfig(configPath,"yml")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	parsedFile, err := ParseConfigFile(configViper)
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	return parsedFile

}

func ParseConfigFile(viper *viper.Viper) (*Config, error) {
	var cfg Config
	err := viper.Unmarshal(&cfg)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &cfg, nil

}

func LoadConfig(filename string,filetype string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType(filetype)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func GetConfigPath(env string) string {
	if env == "development" {
		return "../config/config-dev.yml"
	}else {
		return "../config/config-dev.yml"
	}
	
}
