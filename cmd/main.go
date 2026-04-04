package main

import (
	"co-viewing/internal/app"
	"co-viewing/internal/handler"
	"co-viewing/internal/service"
	"co-viewing/internal/store/postgres"
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	logPath := "logs/app.log"
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		slog.Error("failed to create logs dir", "error", err)
		os.Exit(1)
	}
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		slog.Error("failed to open log file", "path", logPath, "error", err)
		os.Exit(1)
	}
	defer logFile.Close()

	w := io.MultiWriter(os.Stdout, logFile)
	logger := slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	slog.Info("starting subscriptions service")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	slog.Info("loading .env")
	if err := godotenv.Load(); err != nil {
		slog.Error("failed to load .env, using environment", "error", err)
	} else {
		slog.Info("loaded .env")
	}

	conn, err := postgres.NewConnection(os.Getenv("POSTGRES_URL"), ctx)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer conn.Close()
	slog.Info("connected to database")

	userRepo := postgres.NewUserStore(conn)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	srv := &http.Server{
		Addr:    ":9090",
		Handler: app.NewRouter(userHandler),
	}

	go func() {
		slog.Info("server listening", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server listen error", "error", err)
			stop()
		}
	}()

	<-ctx.Done()
	slog.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("server shutdown error", "error", err)
		os.Exit(1)
	}
	slog.Info("server stopped gracefully")
}
