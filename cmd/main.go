package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/mispon/digiseller-express/internal/auth"
	"github.com/mispon/digiseller-express/internal/db"
	"github.com/mispon/digiseller-express/internal/service"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func main() {
	app := gin.Default()
	app.LoadHTMLGlob("templates/*")

	ctx := context.Background()

	token, err := auth.Token()
	if err != nil || token == "" {
		log.Fatal("failed to get token", err)
	}

	provider := mustProvider(ctx)
	defer provider.Close()

	svc := service.New(token, mustLogger(), provider)

	app.GET("/callback", svc.Callback)
	app.GET("/ping", svc.Ping)

	log.Fatal(app.Run())
}

func mustProvider(ctx context.Context) *db.Provider {
	url, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatal("empty DATABASE_URL env")
	}

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	return db.NewProvider(ctx, conn)
}

func mustLogger() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.DisableStacktrace = true

	logLevel := zapcore.InfoLevel
	cfg.Level = zap.NewAtomicLevelAt(logLevel)

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	return logger
}
