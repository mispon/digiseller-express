package service

import (
	"errors"
	"fmt"
	"github.com/mispon/digiseller-express/internal/http"
)

const (
	paymentURL = "https://api.digiseller.ru/api/purchases/unique-code"
	productURL = "https://api.digiseller.ru/api/products/%d/data"
)

type (
	PaymentOpts struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	Payment struct {
		ProductID int           `json:"id_goods"`
		Amount    float64       `json:"amount"`
		Email     string        `json:"email"`
		Method    string        `json:"method"`
		Currency  string        `json:"type_curr"`
		Options   []PaymentOpts `json:"options"`
	}
)

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
		ID      int          `json:"id"`
		Prices  Prices       `json:"prices"`
		Options []ProductOpt `json:"options"`
	}

	Prices struct {
		Initial InitialPrice `json:"initial"`
	}

	InitialPrice struct {
		Rub float64 `json:"RUB"`
	}

	ProductOpt struct {
		Id    int             `json:"id"`
		Label string          `json:"label"`
		Vars  []ProductOptVar `json:"variants"`
	}

	ProductOptVar struct {
		Text        string  `json:"text"`
		ModifyType  string  `json:"modify_type"`
		ModifyValue float64 `json:"modify_value"`
	}
)

func getCodePrice(payment *Payment) (int, error) {
	url := fmt.Sprintf(productURL, payment.ProductID)

	resp, err := http.Do[ProductResp]("GET", url, nil)
	if err != nil {
		return 0, err
	}

	initialPrice := resp.Product.Prices.Initial.Rub

PriceLoop:
	for _, payOpt := range payment.Options {
		for _, prodOpt := range resp.Product.Options {
			if payOpt.Id != prodOpt.Id || payOpt.Name != prodOpt.Label {
				continue
			}

			for _, pov := range prodOpt.Vars {
				if pov.Text != payOpt.Value || pov.ModifyType != "RUB" {
					continue
				}

				initialPrice += pov.ModifyValue
				break PriceLoop
			}
		}
	}

	return int(initialPrice), nil
}
