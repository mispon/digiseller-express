package service

import (
	"github.com/cenk/backoff"
	"github.com/mispon/digiseller-express/internal/auth"
	"github.com/mispon/digiseller-express/internal/db"
	"go.uber.org/zap"
	"time"
)

type Service struct {
	logger   *zap.Logger
	provider *db.Provider
	token    string
}

// New returns new Service
func New(token string, logger *zap.Logger, provider *db.Provider) *Service {
	svc := &Service{
		logger:   logger,
		provider: provider,
		token:    token,
	}

	go svc.refreshToken()

	return svc
}

func (s *Service) refreshToken() {
	ticker := time.NewTicker(100 * time.Minute)
	for {
		<-ticker.C

		delay := backoff.NewExponentialBackOff()
		delay.InitialInterval = 5 * time.Second
		delay.MaxElapsedTime = 0

		for {
			token, err := auth.Token()
			if err != nil {
				s.logger.Error("failed to refresh token, try again", zap.Error(err))
				<-time.After(delay.NextBackOff())
				continue
			}

			s.token = token
			break
		}
	}
}
