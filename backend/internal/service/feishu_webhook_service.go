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

type OpsFeishuAlertCardContext struct {
	EventID int64
	RuleID  int64
	Token   string
}

type feishuWebhookInteractivePayload struct {
	MsgType string            `json:"msg_type"`
	Card    feishuMessageCard `json:"card"`
}

type feishuMessageCard struct {
	Config   feishuMessageCardConfig    `json:"config"`
	Header   feishuMessageCardHeader    `json:"header"`
	Elements []feishuMessageCardElement `json:"elements"`
}

type feishuMessageCardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode"`
}

type feishuMessageCardHeader struct {
	Title    feishuMessageCardText `json:"title"`
	Template string                `json:"template,omitempty"`
}

type feishuMessageCardText struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

type feishuMessageCardElement struct {
	Tag     string                    `json:"tag"`
	Text    *feishuMessageCardText    `json:"text,omitempty"`
	Actions []feishuMessageCardAction `json:"actions,omitempty"`
}

type feishuMessageCardAction struct {
	Tag   string                 `json:"tag"`
	Text  feishuMessageCardText  `json:"text"`
	Type  string                 `json:"type,omitempty"`
	Value map[string]interface{} `json:"value,omitempty"`
}

func sendFeishuWebhookNotification(ctx context.Context, webhookURL string, title string, body string, cardContext *OpsFeishuAlertCardContext) error {
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

	var payload any
	if cardContext != nil && cardContext.EventID > 0 && strings.TrimSpace(cardContext.Token) != "" {
		payload = feishuWebhookInteractivePayload{
			MsgType: "interactive",
			Card:    buildFeishuAlertCard(title, body, cardContext),
		}
	} else {
		payload = feishuWebhookTextPayload{
			MsgType: "text",
			Content: feishuWebhookTextContent{Text: text},
		}
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

func buildFeishuAlertCard(title string, body string, ctx *OpsFeishuAlertCardContext) feishuMessageCard {
	body = strings.TrimSpace(body)
	if body == "" {
		body = strings.TrimSpace(title)
	}
	value := map[string]interface{}{
		"action":   "resolve_alert",
		"event_id": ctx.EventID,
		"rule_id":  ctx.RuleID,
		"token":    strings.TrimSpace(ctx.Token),
	}
	return feishuMessageCard{
		Config: feishuMessageCardConfig{WideScreenMode: true},
		Header: feishuMessageCardHeader{
			Template: "red",
			Title:    feishuMessageCardText{Tag: "plain_text", Content: strings.TrimSpace(title)},
		},
		Elements: []feishuMessageCardElement{
			{
				Tag:  "div",
				Text: &feishuMessageCardText{Tag: "lark_md", Content: body},
			},
			{
				Tag: "action",
				Actions: []feishuMessageCardAction{
					{
						Tag:   "button",
						Text:  feishuMessageCardText{Tag: "plain_text", Content: "标记已处理"},
						Type:  "primary",
						Value: value,
					},
				},
			},
		},
	}
}
