package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type Config struct {
	AppConfig AppConfig
}

type AppConfig struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
	Dsn  string `yaml:"dsn"`
}

func InitConfig() *Config {
	viper.SetConfigName("app-dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	appConfig := &AppConfig{Name: viper.GetString("app.name"), Port: viper.GetInt("app.port"), Dsn: viper.GetString("db.dsn")}

	return &Config{*appConfig}
}
