package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mispon/digiseller-express/internal/http"
)

const (
	paymentURL = "https://api.digiseller.ru/api/purchases/unique-code"
	productURL = "https://api.digiseller.ru/api/products/%d/data"
)

type Payment struct {
	ProductID int     `json:"id_goods"`
	Amount    float64 `json:"amount"`
	Email     string  `json:"email"`
	Method    string  `json:"method"`
	Currency  string  `json:"type_curr"`
}

func getPayment(uniqueCode, token string) (*Payment, error) {
	url := fmt.Sprintf("%s/%s?token=%s", paymentURL, uniqueCode, token)

	payment, err := http.Do[Payment]("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if payment.Amount == 0 && payment.Email == "" {
		return nil, errors.New("failed to parse response")
	}

	return &payment, nil
}

type (
	ProductResp struct {
		Product Product `json:"product"`
	}

	Product struct {
		ID             int             `json:"id"`
		PaymentMethods []PaymentMethod `json:"payment_methods"`
	}

	PaymentMethod struct {
		Code       string     `json:"code"`
		Currencies []Currency `json:"currencies"`
	}

	Currency struct {
		Type  string  `json:"currency"`
		Code  string  `json:"code"`
		Price float64 `json:"price"`
	}
)

func getRubPrice(payment *Payment) (int, error) {
	url := fmt.Sprintf(productURL, payment.ProductID)

	resp, err := http.Do[ProductResp]("GET", url, nil)
	if err != nil {
		return 0, err
	}

	for _, method := range resp.Product.PaymentMethods {
		if !strings.EqualFold(method.Code, payment.Method) {
			// skip other possible methods
			continue
		}

		for _, curr := range method.Currencies {
			if curr.Type == "WMR" || curr.Type == "RUB" {
				return int(curr.Price), nil
			}
		}
	}

	return 0, fmt.Errorf("method %q not found among the possible payment methods", payment.Method)
}
