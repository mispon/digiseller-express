package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Do[T any](method, url string, body io.Reader) (T, error) {
	var empty T
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return empty, fmt.Errorf("failed to create http request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return empty, fmt.Errorf("failed to send http request: %v", err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return empty, fmt.Errorf("failed to read response content: %v", err)
	}

	var result T
	if jErr := json.Unmarshal(content, &result); jErr != nil {
		return empty, fmt.Errorf("failed to unmarshal response: %v", jErr)
	}

	return result, nil
}
