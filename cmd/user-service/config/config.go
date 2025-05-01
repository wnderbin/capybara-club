package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type ServiceConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"service-port"`
	JWTKey  string `yaml:"jwt_key"`
}

func MustLoad() *ServiceConfig {
	config_path := os.Getenv("CONFIG_PATH")

	if config_path == "" {
		log.Fatal("[ config.go ] Config_path is not set")
	}
	if _, err := os.Stat(config_path); os.IsExist(err) {
		log.Fatalf("[ config.go ] Config is not exist: %s\n", config_path)
	}

	var config ServiceConfig

	if err := cleanenv.ReadConfig(config_path, &config); err != nil {
		log.Fatalf("[ config.go ] Cannot read config: %s\n", config_path)
	}
	return &config
}
