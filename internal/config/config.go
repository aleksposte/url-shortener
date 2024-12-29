package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yalm:"env" env:"ENV" env-default:"local"`
	StoragePath string `yalm:"storage_path" env-required:"true"`
	HTTPServer  `yalm:"http_server"`
}

type HTTPServer struct {
	Address     string        `yalm:"address" env-default:":8080"`
	Timeout     time.Duration `yalm:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yalm:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file dosent not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return &cfg
}
