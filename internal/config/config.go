package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `json:"env" env-required:"true"`
	Storage    string     `json:"storage_path" env-required:"true"`
	HTTPServer HTTPServer `json:"http_server"`
}

type HTTPServer struct {
	Address     string            `json:"address" env-required:"true"`
	Timeout     time.Duration `json:"timeout" env-required:"false"`
	IdleTimeout time.Duration `json:"idle_timeout" env-required:"false"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file is not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	
	return &cfg
}
