package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestUpdateEmailNotificationConfig_PersistsFeishuWebhook(t *testing.T) {
	repo := newRuntimeSettingRepoStub()
	svc := &OpsService{settingRepo: repo}

	updated, err := svc.UpdateEmailNotificationConfig(context.Background(), &OpsEmailNotificationConfigUpdateRequest{
		Alert: &OpsEmailAlertConfig{
			Enabled:               true,
			Recipients:            []string{"admin@example.com"},
			MinSeverity:           "warning",
			RateLimitPerHour:      10,
			BatchingWindowSeconds: 0,
			IncludeResolvedAlerts: false,
		},
		Report: &OpsEmailReportConfig{
			Enabled:                         false,
			Recipients:                      []string{},
			DailySummaryEnabled:             false,
			DailySummarySchedule:            "0 9 * * *",
			WeeklySummaryEnabled:            false,
			WeeklySummarySchedule:           "0 9 * * 1",
			ErrorDigestEnabled:              false,
			ErrorDigestSchedule:             "0 9 * * *",
			ErrorDigestMinCount:             10,
			AccountHealthEnabled:            false,
			AccountHealthSchedule:           "0 9 * * *",
			AccountHealthErrorRateThreshold: 10,
		},
		Feishu: &OpsFeishuNotificationConfig{
			Alert: OpsFeishuAlertConfig{
				Enabled:     true,
				WebhookURLs: []string{" https://open.feishu.cn/open-apis/bot/v2/hook/test-token "},
			},
		},
	})
	if err != nil {
		t.Fatalf("UpdateEmailNotificationConfig() error = %v", err)
	}
	if !updated.Feishu.Alert.Enabled {
		t.Fatalf("Feishu alert enabled = false, want true")
	}
	if got := updated.Feishu.Alert.WebhookURLs; len(got) != 1 || got[0] != "https://open.feishu.cn/open-apis/bot/v2/hook/test-token" {
		t.Fatalf("Feishu webhook urls = %#v", got)
	}

	reloaded, err := svc.GetEmailNotificationConfig(context.Background())
	if err != nil {
		t.Fatalf("GetEmailNotificationConfig() error = %v", err)
	}
	if got := reloaded.Feishu.Alert.WebhookURLs; len(got) != 1 || got[0] != "https://open.feishu.cn/open-apis/bot/v2/hook/test-token" {
		t.Fatalf("reloaded Feishu webhook urls = %#v", got)
	}
}

func TestSendFeishuWebhookNotification_PostsTextPayload(t *testing.T) {
	var gotContentType string
	var gotPayload map[string]any
	client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		gotContentType = r.Header.Get("Content-Type")
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if err := json.NewDecoder(r.Body).Decode(&gotPayload); err != nil {
			t.Fatalf("decode request body: %v", err)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{"StatusCode":0,"msg":"success"}`)),
		}, nil
	})}

	oldClient := feishuWebhookHTTPClient
	feishuWebhookHTTPClient = client
	defer func() {
		feishuWebhookHTTPClient = oldClient
	}()

	err := sendFeishuWebhookNotification(context.Background(), "https://open.feishu.cn/open-apis/bot/v2/hook/test", "Ops Alert", "service unavailable")
	if err != nil {
		t.Fatalf("sendFeishuWebhookNotification() error = %v", err)
	}
	if !strings.HasPrefix(gotContentType, "application/json") {
		t.Fatalf("Content-Type = %q, want application/json", gotContentType)
	}
	if gotPayload["msg_type"] != "text" {
		t.Fatalf("msg_type = %#v, want text", gotPayload["msg_type"])
	}
	content, ok := gotPayload["content"].(map[string]any)
	if !ok {
		t.Fatalf("content = %#v, want object", gotPayload["content"])
	}
	text, _ := content["text"].(string)
	if !strings.Contains(text, "Ops Alert") || !strings.Contains(text, "service unavailable") {
		t.Fatalf("text payload = %q", text)
	}
}

func TestMaybeSendAlertEmail_SendsFeishuWebhookWhenConfigured(t *testing.T) {
	var webhookCalls int
	client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		webhookCalls++
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{"StatusCode":0,"msg":"success"}`)),
		}, nil
	})}

	oldClient := feishuWebhookHTTPClient
	feishuWebhookHTTPClient = client
	defer func() {
		feishuWebhookHTTPClient = oldClient
	}()

	repo := newRuntimeSettingRepoStub()
	cfg := defaultOpsEmailNotificationConfig()
	cfg.Alert.Enabled = false
	cfg.Feishu.Alert.Enabled = true
	cfg.Feishu.Alert.WebhookURLs = []string{"https://open.feishu.cn/open-apis/bot/v2/hook/test"}
	raw, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal config: %v", err)
	}
	repo.values[SettingKeyOpsEmailNotificationConfig] = string(raw)

	opsSvc := &OpsService{settingRepo: repo}
	evaluator := NewOpsAlertEvaluatorService(opsSvc, &opsRepoMock{}, nil, nil, nil)
	rule := &OpsAlertRule{ID: 1, Name: "High error rate", Severity: "critical", MetricType: "error_rate", Operator: ">", Threshold: 5, NotifyEmail: true}
	value := 9.2
	event := &OpsAlertEvent{ID: 10, RuleID: 1, Status: "firing", MetricValue: &value, FiredAt: time.Now().UTC(), Description: "error_rate above threshold"}

	if sent := evaluator.maybeSendAlertEmail(context.Background(), nil, rule, event); !sent {
		t.Fatalf("maybeSendAlertEmail() = false, want true")
	}
	if webhookCalls != 1 {
		t.Fatalf("webhook calls = %d, want 1", webhookCalls)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
