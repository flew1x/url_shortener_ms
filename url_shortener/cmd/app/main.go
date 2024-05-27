package main

import (
	"context"

	"github.com/flew1x/url_shortener_ms/internal/app"
	"github.com/flew1x/url_shortener_ms/internal/config"
	"github.com/flew1x/url_shortener_ms/pkg/logger"
)

func main() {
	ctx := context.Background()

	cfg := config.NewConfig()
	cfg.InitConfig(CONFIG_PATH_ENV, CONFIG_FILE_ENV)

	logger := logger.InitLogger(cfg.GetLoggerConfig().GetLogLevel())

	server, err := app.InitialServer(ctx, cfg, logger)
	if err != nil {
		panic(err)
	}

	server.Run()
}
