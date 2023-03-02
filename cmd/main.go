package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mispon/digiseller-express/internal/db"
	"github.com/mispon/digiseller-express/internal/service"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func main() {
	app := gin.Default()
	app.LoadHTMLGlob("templates/*")

	provider := db.NewProvider()
	svc := service.New(mustLogger(), provider)

	app.GET("/callback", svc.Callback)
	app.GET("/ping", svc.Ping)

	log.Fatal(app.Run())
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
