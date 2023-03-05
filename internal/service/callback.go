package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/cenk/backoff"
	"github.com/gin-gonic/gin"
	"github.com/mispon/digiseller-express/internal/http"
	"go.uber.org/zap"
)

const checkPaymentURL = "https://api.digiseller.ru/api/purchases/unique-code"

type Payment struct {
	Amount float64 `json:"amount"`
	Email  string  `json:"email"`
}

func (s *Service) Callback(c *gin.Context) {
	uniqueCode, ok := c.GetQuery("uniquecode")
	if !ok {
		c.JSON(200, gin.H{"status": "ok"})
		return
	}

	if uniqueCode == "" {
		c.HTML(400, "error.tmpl", gin.H{"message": "400 - пустой uniquecode"})
		return
	}

	payment, err := getPayment(uniqueCode, s.token)
	if err != nil {
		s.logger.Error("failed to check payment", zap.Error(err))
		c.HTML(502, "error.tmpl", gin.H{"message": "502 - ошибка проверки платежа"})
		return
	}

	issuedCode, ok := s.provider.GetIssued(uniqueCode)
	if ok {
		c.HTML(200, "index.tmpl", gin.H{"code": issuedCode})
		return
	}

	price := int(payment.Amount)

	code, err := s.provider.PopCode(price)
	if err != nil {
		s.logger.Error("failed to get code", zap.Error(err))
		c.HTML(502, "error.tmpl", gin.H{"message": "502 - отсутствует подходящий код оплаты"})
		return
	}

	go s.saveIssuedCode(uniqueCode, code, price, payment.Email)

	c.HTML(200, "index.tmpl", gin.H{"code": code})
}

func getPayment(uniqueCode, token string) (Payment, error) {
	url := fmt.Sprintf("%s/%s?token=%s", checkPaymentURL, uniqueCode, token)

	payment, err := http.Do[Payment]("GET", url, nil)
	if err != nil {
		return Payment{}, err
	}

	if payment.Amount == 0 && payment.Email == "" {
		return Payment{}, errors.New("failed to parse response")
	}

	return payment, nil
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
