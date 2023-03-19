package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/mispon/digiseller-express/internal/auth"
	"github.com/mispon/digiseller-express/internal/db"
	"github.com/mispon/digiseller-express/internal/env"
	"github.com/mispon/digiseller-express/internal/service"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var (
	templateFolders = []string{"html", "templates"}
)

func main() {
	files, err := readHTMLFiles()
	if err != nil {
		log.Fatal("failed to read html templates", err)
	}

	app := gin.Default()
	app.LoadHTMLFiles(files...)

	ctx := context.Background()

	token, err := auth.Token()
	if err != nil || token == "" {
		log.Fatal("failed to get token", err)
	}

	provider := mustProvider(ctx)
	defer provider.Close()

	svc := service.New(token, mustLogger(), provider)

	app.GET("/callback", svc.Callback)
	app.POST("/callback", svc.Callback)
	app.GET("/ping", svc.Ping)

	log.Fatal(app.Run())
}

func mustProvider(ctx context.Context) *db.Provider {
	url := env.DatabaseURL()
	if url == "" {
		log.Fatal("empty database url")
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

func readHTMLFiles() ([]string, error) {
	filesMap := make(map[string]string)

	for _, folder := range templateFolders {
		entries, err := os.ReadDir(folder)
		if err != nil {
			return nil, err
		}

		for _, entry := range entries {
			fileName := entry.Name()
			if _, ok := filesMap[fileName]; !ok {
				filesMap[fileName] = folder + "/" + fileName
			}
		}
	}

	files := make([]string, 0, len(filesMap))
	for _, path := range filesMap {
		files = append(files, path)
	}

	return files, nil
}
