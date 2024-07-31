package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cvckeboy/restaurant-app/utils/config"
	"github.com/fatih/color"

	"io"
	"log"
	"log/slog"
	"os"
)

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(timeStr, level, msg, color.WhiteString(string(b)))

	return nil
}

func MyLogger(cfg *config.Config) *PrettyHandler {
	var programLevel = new(slog.LevelVar)
	switch cfg.Server.Level {
	case "debug":
		programLevel.Set(slog.LevelDebug)
	case "info":
		programLevel.Set(slog.LevelInfo)
	case "warn":
		programLevel.Set(slog.LevelWarn)
	case "error":
		programLevel.Set(slog.LevelError)
	}

	file, err := os.OpenFile("data", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	out := io.MultiWriter(os.Stdout, file)

	lg := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &slog.HandlerOptions{
			Level: programLevel,
		}),
		l: log.New(out, "", 0),
	}
	return lg
}
