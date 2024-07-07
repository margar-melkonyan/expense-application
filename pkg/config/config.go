package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	HttpServer `yaml:"http_server"`
	DB         `yaml:"db"`
}

type HttpServer struct {
	Port        string        `yaml:"port" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"10s"`
}

type DB struct {
	Host    string `yaml:"host" env-required:"true"`
	Name    string `yaml:"name" env-required:"true"`
	User    string `yaml:"username" env-required:"true"`
	SSLMode string `yaml:"ssl_mode"`
	Port    string `yaml:"port" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		slog.Error("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Panic(fmt.Sprintf("Config file doesn't exists: %s", configPath))
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		slog.Error(fmt.Sprintf("Cannot read config %s", err))
	}

	return &cfg
}
