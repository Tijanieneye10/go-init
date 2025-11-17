package exceptions

import (
	"fmt"
	"log/slog"
	"os"
)

type CustomError struct {
	Msg string
}

func (c *CustomError) Error() string {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		slog.Error("Failed to open log file", "error", err)
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(file, nil))

	// Logs will be in JSON format
	logger.Info(
		"User logged in",
		"user_id", 12325,
		"ip", "192",
	)
	return fmt.Sprintf("%v", c.Msg)
}
