package env

import (
	"os"
	"strings"
)

const (
	databaseUrl  = "DATABASE_URL"
	sellerId     = "SELLER_ID"
	sellerApiKey = "SELLER_API_KEY"
	tgUser       = "TG_USER"
)

func DatabaseURL() string {
	return os.Getenv(databaseUrl)
}

func SellerId() string {
	return os.Getenv(sellerId)
}

func SellerApiKey() string {
	return os.Getenv(sellerApiKey)
}

func TelegramUser() string {
	val := os.Getenv(tgUser)
	return strings.TrimPrefix(val, "@")
}
