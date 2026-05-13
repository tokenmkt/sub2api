package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var feishuWebhookHTTPClient = &http.Client{Timeout: 10 * time.Second}

type feishuWebhookTextPayload struct {
	MsgType string                   `json:"msg_type"`
	Content feishuWebhookTextContent `json:"content"`
}

type feishuWebhookTextContent struct {
	Text string `json:"text"`
}

func sendFeishuWebhookNotification(ctx context.Context, webhookURL string, title string, body string) error {
	webhookURL = strings.TrimSpace(webhookURL)
	if webhookURL == "" {
		return fmt.Errorf("feishu webhook url is required")
	}
	if ctx == nil {
		ctx = context.Background()
	}

	text := strings.TrimSpace(title)
	if content := strings.TrimSpace(body); content != "" {
		if text != "" {
			text += "\n\n"
		}
		text += content
	}
	if text == "" {
		return fmt.Errorf("feishu message text is required")
	}

	payload := feishuWebhookTextPayload{
		MsgType: "text",
		Content: feishuWebhookTextContent{Text: text},
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, webhookURL, bytes.NewReader(raw))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := feishuWebhookHTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("feishu webhook returned HTTP %d: %s", resp.StatusCode, strings.TrimSpace(string(respBody)))
	}

	var result struct {
		StatusCode *int   `json:"StatusCode"`
		Code       *int   `json:"code"`
		Msg        string `json:"msg"`
	}
	if len(respBody) > 0 && json.Unmarshal(respBody, &result) == nil {
		code := 0
		if result.StatusCode != nil {
			code = *result.StatusCode
		} else if result.Code != nil {
			code = *result.Code
		}
		if code != 0 {
			return fmt.Errorf("feishu webhook returned code %d: %s", code, strings.TrimSpace(result.Msg))
		}
	}
	return nil
}
