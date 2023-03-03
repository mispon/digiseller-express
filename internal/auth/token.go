package auth

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mispon/digiseller-express/internal/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const tokenURL = "https://api.digiseller.ru/api/apilogin"

type TokenWrap struct {
	Value string `json:"token"`
}

// Token creates new auth token for digiseller API
func Token() (string, error) {
	sellerIdStr, ok := os.LookupEnv("SELLER_ID")
	if !ok {
		return "", errors.New("empty SELLER_ID env")
	}

	sellerId, err := strconv.Atoi(sellerIdStr)
	if err != nil {
		return "", err
	}

	sellerApiKey, ok := os.LookupEnv("SELLER_API_KEY")
	if !ok {
		return "", errors.New("empty SELLER_API_KEY env")
	}

	ts := time.Now().Unix()
	sign := createSign(sellerApiKey, ts)

	request := map[string]any{
		"seller_id": sellerId,
		"timestamp": ts,
		"sign":      sign,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	token, err := http.Do[TokenWrap]("POST", tokenURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	return token.Value, nil
}

func createSign(apiKey string, ts int64) string {
	data := fmt.Sprintf("%s%d", apiKey, ts)

	s := sha256.New()
	s.Write([]byte(data))

	return strings.TrimRight(fmt.Sprintf("%x\n", s.Sum(nil)), "\n")
}
