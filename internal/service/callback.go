package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const checkPaymentURL = "https://api.digiseller.ru/api/purchases/unique-code/"

type Payment struct {
	Amount int
	Email  string
	Date   string
	Status string
}

func (s *Service) Callback(c *gin.Context) {
	uniqueCode := c.Query("uniquecode")
	if uniqueCode == "" {
		c.JSON(400, gin.H{"message": "empty unique code"})
		return
	}

	payment, err := getPayment(uniqueCode)
	if err != nil {
		s.logger.Error("failed to check payment", zap.Error(err))
		c.JSON(502, gin.H{
			"message": fmt.Errorf("failed to check payment: %v", err),
		})
		return
	}

	code, err := s.provider.PopCode(payment.Amount)
	if err != nil {
		s.logger.Error("failed to check payment", zap.Error(err))
		c.JSON(502, gin.H{
			"message": fmt.Errorf("failed to get code: %v", err),
		})
		return
	}

	if err = s.provider.SaveIssued(uniqueCode, code, payment.Amount, payment.Email, payment.Date); err != nil {
		s.logger.Error("failed to save issued code",
			zap.String("unique_code", uniqueCode),
			zap.String("code", code),
			zap.Any("payment", payment),
			zap.Error(err),
		)
	}

	c.HTML(200, "index.tmpl", gin.H{
		"title": "Your order is confirmed!",
		"code":  code,
	})
}

func getPayment(uniqueCode string) (Payment, error) {
	url := fmt.Sprintf("%s/%s", checkPaymentURL, uniqueCode)
	fmt.Println(url)
	return Payment{
		Amount: 20,
		Email:  "pupa@lupa.com",
		Date:   "now!",
		Status: "ok",
	}, nil
}
