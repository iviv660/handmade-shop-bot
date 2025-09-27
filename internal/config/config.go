package config

import (
	"log"
	"os"
	"strconv"
)

type config struct {
	PostgresDSN string
	TokenTG     string
	ShopID      string
	SecretKey   string
	BotURL      string
	AdminID     int64
}

var C config

func Load() {
	admin := getEnv("ADMIN_ID", "0")
	adminID, err := strconv.ParseInt(admin, 10, 64)
	if err != nil {
		log.Fatalf("❌ invalid ADMIN_ID: %v", err)
	}

	C = config{
		PostgresDSN: getEnv("POSTGRES_DSN", ""),
		TokenTG:     getEnv("TELEGRAM_TOKEN", ""),
		ShopID:      getEnv("SHOP_ID", ""),
		SecretKey:   getEnv("SECRET_KEY", ""),
		BotURL:      getEnv("BOT_URL", ""),
		AdminID:     adminID,
	}

	if C.PostgresDSN == "" {
		log.Fatal("❌ POSTGRES_DSN is required")
	}
	if C.TokenTG == "" {
		log.Fatal("❌ TELEGRAM_TOKEN is required")
	}

	log.Printf("✅ Config loaded (DB=%s, BOT_URL=%s, ADMIN_ID=%d)", C.PostgresDSN, C.BotURL, C.AdminID)
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
