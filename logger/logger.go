package logger

import (
	"log/slog"
	"os"
)

func Init(prod bool) {
	var handler slog.Handler
	if prod {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})
	}

	slog.SetDefault(slog.New(handler))
}
