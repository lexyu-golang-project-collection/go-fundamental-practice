package main

import (
	"log/slog"
	"os"
)

/*
2024/05/29 11:05:18 INFO hello, world user=""
time=2024-05-29T11:05:18.867+08:00 level=INFO msg="hello, world" user=""
*/
func main() {
	logger := slog.Default()
	logger.Info("hello, world", "user", os.Getenv("USER"))

	logger2 := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger2.Info("hello, world", "user", os.Getenv("USER"))
}
