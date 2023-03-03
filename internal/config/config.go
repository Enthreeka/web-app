package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type config struct {
	Server struct {
		Hostname   string `yaml:"hostname"`
		TypeServer string `yaml:"type"`
		Port       string `yaml:"port"`
	} `yaml:"server"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	DatabaseName string `yaml:"database_name"`
	Url          string `yaml:"url"`
}

var instance *config

func GetConfig() *config {
	log.Println("Сonfig is running...")

	cfg := viper.New()
	cfg.SetConfigType("yaml")
	cfg.AddConfigPath("configs")
	cfg.SetConfigName("config")
	err := cfg.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	instance = &config{}

	err = cfg.Unmarshal(instance)
	if err != nil {
		log.Println(err)
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return instance
}
