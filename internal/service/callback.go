package service

import (
	"fmt"
	"time"

	"github.com/cenk/backoff"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const checkPaymentURL = "https://api.digiseller.ru/api/purchases/unique-code/"

type Payment struct {
	Amount int
	Email  string
	Status string
}

func (s *Service) Callback(c *gin.Context) {
	uniqueCode := c.Query("uniquecode")
	if uniqueCode == "" {
		c.HTML(400, "error.tmpl", gin.H{"message": "распознать uniquecode"})
		return
	}

	payment, err := getPayment(uniqueCode)
	if err != nil {
		s.logger.Error("failed to check payment", zap.Error(err))
		c.HTML(502, "error.tmpl", gin.H{
			"action": "проверить платеж",
		})
		return
	}

	code, err := s.provider.PopCode(payment.Amount)
	if err != nil {
		s.logger.Error("failed to get code", zap.Error(err))
		c.HTML(502, "error.tmpl", gin.H{
			"action": "получить код оплаты",
		})
		return
	}

	c.HTML(200, "index.tmpl", gin.H{
		"title": "Your order is confirmed!",
		"code":  code,
	})
}

func getPayment(uniqueCode string) (Payment, error) {
	url := fmt.Sprintf("%s/%s", checkPaymentURL, uniqueCode)
	fmt.Println(url)

	// todo

	return Payment{
		Amount: 20,
		Email:  "pupa@lupa.com",
		Status: "ok",
	}, nil
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
