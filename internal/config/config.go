package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `json:"env" env-required:"true"`
	Storage    string `json:"storage_path" env-required:"true"`
	HTTPServer `json:"http_server"`
}

type HTTPServer struct {
	Address     string        `json:"address" env-required:"true"`
	Timeout     time.Duration `json:"timeout" env-required:"false"`
	IdleTimeout time.Duration `json:"idle_timeout" env-required:"false"`
}

func MustLoad() *Config{
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}

	return &cfg
}
