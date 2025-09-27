package app

import (
	"app/internal/bot"
	"app/internal/config"
	"app/internal/logger"
	"app/internal/service"
	"app/internal/storage"
	"context"
	"fmt"
	"os"
	"time"

	"gopkg.in/telebot.v4"
)

func Run() error {
	config.Load()

	gorm, err := storage.ConnectPostgres(context.Background(), config.C.PostgresDSN)
	if err != nil {
		return fmt.Errorf("connect postgres: %w", err)
	}

	repoUs := storage.NewUserDB(gorm)
	repoPr := storage.NewProductDB(gorm)
	repoOr := storage.NewOrderDB(gorm)

	logg := logger.New()
	uc := service.NewUseCase(repoUs, repoPr, repoOr, nil, logg)

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		return fmt.Errorf("TELEGRAM_TOKEN is not set")
	}
	b, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return fmt.Errorf("create app: %w", err)
	}

	bot.RegisterHandlers(b, uc, config.C.AdminID)

	b.Start()
	return nil
}
