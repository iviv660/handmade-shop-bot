package youkassa

import (
	"app/internal/dto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type YooKassa struct {
	shopID     string
	secretKey  string
	botURL     string
	httpClient *http.Client
}

func NewYooKassa(shopID, secretKey, botURL string) *YooKassa {
	return &YooKassa{
		shopID:    shopID,
		secretKey: secretKey,
		botURL:    botURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (y *YooKassa) Create(ctx context.Context, order *dto.Order) (string, string, error) {
	requestBody := map[string]any{
		"amount": map[string]string{
			"value":    fmt.Sprintf("%.2f", order.TotalPrice),
			"currency": "RUB",
		},
		"capture": true,
		"confirmation": map[string]string{
			"type":       "redirect",
			"return_url": y.botURL,
		},
		"description": fmt.Sprintf("Оплата заказа №%d", order.ID),
		"metadata": map[string]any{
			"order_id": order.ID,
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return "", "", fmt.Errorf("marshal youkassa body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.yookassa.ru/v3/payments", bytes.NewBuffer(body))
	if err != nil {
		return "", "", fmt.Errorf("create request: %w", err)
	}
	req.SetBasicAuth(y.shopID, y.secretKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := y.httpClient.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", "", fmt.Errorf("yookassa returned status %s", resp.Status)
	}

	var res struct {
		ID           string `json:"id"`
		Confirmation struct {
			URL string `json:"confirmation_url"`
		} `json:"confirmation"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", "", fmt.Errorf("decode response: %w", err)
	}

	return res.Confirmation.URL, res.ID, nil
}

func (y *YooKassa) CheckStatus(ctx context.Context, paymentID string) (string, error) {
	url := fmt.Sprintf("https://api.yookassa.ru/v3/payments/%s", paymentID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.SetBasicAuth(y.shopID, y.secretKey)

	resp, err := y.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("yookassa returned status %s", resp.Status)
	}

	var res struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	return res.Status, nil
}
