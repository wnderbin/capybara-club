package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSL_Mode string `yaml:"sslmode"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type DBConfig struct {
	Env           string         `yaml:"env"`
	Postgres      PostgresConfig `yaml:"postgres"`
	Redis         RedisConfig    `yaml:"redis"`
	StartUpStatus int            `yaml:"startup-status"`
}

func MustLoad() *DBConfig {
	config_path := os.Getenv("CONFIG_PATH")

	if config_path == "" {
		log.Fatal("[ config.go ] Config_path is not set")
	}
	if _, err := os.Stat(config_path); os.IsExist(err) {
		log.Fatalf("[ config.go ] Config is not exist: %s\n", config_path)
	}

	var config DBConfig

	if err := cleanenv.ReadConfig(config_path, &config); err != nil {
		log.Fatalf("[ config.go ] Cannot read config: %s\n", config_path)
	}
	return &config
}

func DatabaseLoad() *DBConfig {
	config_path := "./internal/config/config.yaml"

	if config_path == "" {
		log.Fatal("[ config.go ] Config_path is not set")
	}
	if _, err := os.Stat(config_path); os.IsExist(err) {
		log.Fatalf("[ config.go ] Config is not exist: %s\n", config_path)
	}

	var config DBConfig

	if err := cleanenv.ReadConfig(config_path, &config); err != nil {
		log.Fatalf("[ config.go ] Cannot read config: %s\n", config_path)
	}
	return &config
}

func (p PostgresConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.DBName, p.SSL_Mode,
	)
}
