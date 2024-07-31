package main

import (
	"github.com/cvckeboy/restaurant-app/utils/config"
	"github.com/cvckeboy/restaurant-app/utils/logger"
	"log/slog"
)

func main() {
	cfg, err := config.NewConfig("config.yaml")
	if err != nil {
		return
	}
	handler := logger.MyLogger(cfg)
	lg := slog.New(handler)
	lg.Info("Config loaded, logger set", "program level", cfg.Server.Level)

}
