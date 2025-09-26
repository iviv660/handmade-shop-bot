package config

import (
	"log"
	"os"
)

type config struct {
	Postgres_DSN string
}

var C config

func Load() {
	C = config{
		Postgres_DSN: os.Getenv("POSTGRES_DSN"),
	}
	log.Println("✅Config loaded")
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
