package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func cmdWebhookSet(url string) {
	token := getToken()
	resp := apiCall(token, "setWebhook", map[string]string{"url": url})

	if resp.Ok {
		fmt.Printf("Webhook set to: %s\n", url)
	} else {
		fatal(fmt.Sprintf("Failed to set webhook: %s (code %d)", resp.Description, resp.ErrorCode))
	}
}

func cmdWebhookDelete() {
	token := getToken()
	resp := apiCall(token, "deleteWebhook", nil)

	if resp.Ok {
		fmt.Println("Webhook deleted.")
	} else {
		fatal(fmt.Sprintf("Failed to delete webhook: %s (code %d)", resp.Description, resp.ErrorCode))
	}
}

func cmdWebhookInfo() {
	token := getToken()
	resp := apiCall(token, "getWebhookInfo", nil)

	if !resp.Ok {
		fatal(fmt.Sprintf("Failed to get webhook info: %s", resp.Description))
	}

	var info map[string]any
	if err := json.Unmarshal(resp.Result, &info); err != nil {
		fatal(fmt.Sprintf("Failed to parse webhook info: %v", err))
	}

	fmt.Println("Webhook Info:")
	for k, v := range info {
		fmt.Printf("  %-30s %v\n", k+":", v)
	}
}

type apiResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"`
	Description string          `json:"description,omitempty"`
	ErrorCode   int             `json:"error_code,omitempty"`
}

func apiCall(token, method string, params map[string]string) apiResponse {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, method)

	var body []byte
	if params != nil {
		body, _ = json.Marshal(params)
	} else {
		body = []byte("{}")
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fatal(fmt.Sprintf("Request failed: %v", err))
	}
	defer resp.Body.Close()

	var result apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fatal(fmt.Sprintf("Failed to decode response: %v", err))
	}
	return result
}

func getToken() string {
	loadEnv()
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		fatal("BOT_TOKEN is not set. Set it in .env or as environment variable.")
	}
	return token
}
