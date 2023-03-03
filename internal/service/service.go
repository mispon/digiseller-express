package service

import (
	"github.com/mispon/digiseller-express/internal/db"
	"go.uber.org/zap"
)

type Service struct {
	logger   *zap.Logger
	provider *db.Provider
}

// New returns new Service
func New(logger *zap.Logger, provider *db.Provider) *Service {
	return &Service{
		logger:   logger.Named("service"),
		provider: provider,
	}
}
