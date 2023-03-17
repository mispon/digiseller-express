package service

import (
	"time"

	"github.com/cenk/backoff"
	"github.com/gin-gonic/gin"
	"github.com/mispon/digiseller-express/internal/env"
	"go.uber.org/zap"
)

func (s *Service) Callback(c *gin.Context) {
	uniqueCode, ok := c.GetQuery("uniquecode")
	if !ok {
		c.JSON(200, gin.H{"status": "ok"})
		return
	}

	if uniqueCode == "" {
		c.HTML(400, "error.tmpl", gin.H{
			"message":  "400 - пустой uniquecode",
			"sellerId": env.SellerId(),
			"tgUser":   env.TelegramUser(),
		})
		return
	}

	issuedCode, ok := s.provider.GetIssued(uniqueCode)
	if ok {
		c.HTML(200, "index.tmpl", gin.H{
			"code":     issuedCode,
			"sellerId": env.SellerId(),
			"tgUser":   env.TelegramUser(),
		})
		return
	}

	payment, err := getPayment(uniqueCode, s.token)
	if err != nil {
		s.logger.Error("failed to check payment", zap.Error(err))
		c.HTML(502, "error.tmpl", gin.H{
			"message":  "502 - ошибка проверки платежа",
			"sellerId": env.SellerId(),
			"tgUser":   env.TelegramUser(),
		})
		return
	}

	price, err := getRubPrice(payment)
	if err != nil {
		s.logger.Error("failed to get rub price", zap.Error(err))
		c.HTML(502, "error.tmpl", gin.H{
			"message":  "502 - ошибка получения стоимости в RUB",
			"sellerId": env.SellerId(),
			"tgUser":   env.TelegramUser(),
		})
		return
	}

	code, err := s.provider.PopCode(payment.ProductID, price)
	if err != nil {
		s.logger.Error("failed to get code", zap.Error(err))
		c.HTML(502, "error.tmpl", gin.H{
			"message":  "502 - отсутствует подходящий код оплаты",
			"sellerId": env.SellerId(),
			"tgUser":   env.TelegramUser(),
		})
		return
	}

	go s.saveIssuedCode(uniqueCode, code, price, payment.Email)

	c.HTML(200, "index.tmpl", gin.H{
		"code":     code,
		"sellerId": env.SellerId(),
		"tgUser":   env.TelegramUser(),
	})
}

func (s *Service) saveIssuedCode(uniqueCode string, code string, price int, email string) {
	delay := backoff.NewExponentialBackOff()

	var err error
	for {
		if err = s.provider.SaveIssued(uniqueCode, code, price, email); err == nil {
			// happy path
			return
		}

		nextBo := delay.NextBackOff()
		if nextBo == backoff.Stop {
			break
		}

		<-time.After(nextBo)
	}

	// if backoff limit is over, just log data
	s.logger.Error("failed to save issued code",
		zap.String("unique_code", uniqueCode),
		zap.String("code", code),
		zap.Any("payment", price),
		zap.String("email", email),
		zap.Error(err),
	)
}
