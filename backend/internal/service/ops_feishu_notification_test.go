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

func TestSendFeishuWebhookNotification_PostsInteractiveCardPayload(t *testing.T) {
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

	err := sendFeishuWebhookNotification(context.Background(), "https://open.feishu.cn/open-apis/bot/v2/hook/test", "Ops Alert", "service unavailable", &OpsFeishuAlertCardContext{
		EventID: 10,
		RuleID:  1,
		Token:   "card-action-token",
	})
	if err != nil {
		t.Fatalf("sendFeishuWebhookNotification() error = %v", err)
	}
	if !strings.HasPrefix(gotContentType, "application/json") {
		t.Fatalf("Content-Type = %q, want application/json", gotContentType)
	}
	if gotPayload["msg_type"] != "interactive" {
		t.Fatalf("msg_type = %#v, want interactive", gotPayload["msg_type"])
	}
	card, ok := gotPayload["card"].(map[string]any)
	if !ok {
		t.Fatalf("card = %#v, want object", gotPayload["card"])
	}
	rawCard, err := json.Marshal(card)
	if err != nil {
		t.Fatalf("marshal card: %v", err)
	}
	cardText := string(rawCard)
	for _, want := range []string{"Ops Alert", "service unavailable", "resolve_alert", "card-action-token", `"event_id":10`} {
		if !strings.Contains(cardText, want) {
			t.Fatalf("card payload missing %q: %s", want, cardText)
		}
	}

	elements, ok := card["elements"].([]any)
	if !ok || len(elements) < 1 {
		t.Fatalf("card elements = %#v, want at least one content element", card["elements"])
	}
	contentElement, ok := elements[0].(map[string]any)
	if !ok {
		t.Fatalf("first card element = %#v, want object", elements[0])
	}
	if got := contentElement["tag"]; got != "div" {
		t.Fatalf("first card element tag = %#v, want div so Feishu renders lark_md text", got)
	}
	textElement, ok := contentElement["text"].(map[string]any)
	if !ok {
		t.Fatalf("first card element text = %#v, want object", contentElement["text"])
	}
	if got := textElement["tag"]; got != "lark_md" {
		t.Fatalf("first card text tag = %#v, want lark_md", got)
	}
	if got := textElement["content"]; got != "service unavailable" {
		t.Fatalf("first card text content = %#v, want service unavailable", got)
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
	cfg.Feishu.Alert.ActionToken = "card-action-token"
	raw, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal config: %v", err)
	}
	repo.values[SettingKeyOpsEmailNotificationConfig] = string(raw)

	opsSvc := &OpsService{settingRepo: repo}
	evaluator := NewOpsAlertEvaluatorService(opsSvc, &opsRepoMock{}, nil, nil, nil)
	rule := &OpsAlertRule{ID: 1, Name: "High error rate", Severity: "critical", MetricType: "error_rate", Operator: ">", Threshold: 5, NotifyEmail: true, NotifyFeishu: true}
	value := 9.2
	event := &OpsAlertEvent{ID: 10, RuleID: 1, Status: "firing", MetricValue: &value, FiredAt: time.Now().UTC(), Description: "error_rate above threshold"}

	if sent := evaluator.maybeSendAlertEmail(context.Background(), nil, rule, event); !sent {
		t.Fatalf("maybeSendAlertEmail() = false, want true")
	}
	if webhookCalls != 1 {
		t.Fatalf("webhook calls = %d, want 1", webhookCalls)
	}
}

func TestMaybeSendAlertEmail_SendsFeishuWhenEmailDisabledForRule(t *testing.T) {
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
	cfg.Feishu.Alert.ActionToken = "card-action-token"
	raw, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal config: %v", err)
	}
	repo.values[SettingKeyOpsEmailNotificationConfig] = string(raw)

	opsSvc := &OpsService{settingRepo: repo}
	evaluator := NewOpsAlertEvaluatorService(opsSvc, &opsRepoMock{}, nil, nil, nil)
	rule := &OpsAlertRule{ID: 1, Name: "High error rate", Severity: "critical", MetricType: "error_rate", Operator: ">", Threshold: 5, NotifyEmail: false, NotifyFeishu: true}
	value := 9.2
	event := &OpsAlertEvent{ID: 10, RuleID: 1, Status: "firing", MetricValue: &value, FiredAt: time.Now().UTC(), Description: "error_rate above threshold"}

	if sent := evaluator.maybeSendAlertEmail(context.Background(), nil, rule, event); !sent {
		t.Fatalf("maybeSendAlertEmail() = false, want true")
	}
	if webhookCalls != 1 {
		t.Fatalf("webhook calls = %d, want 1", webhookCalls)
	}
}

func TestMaybeSendAlertEmail_DoesNotSendFeishuWhenRuleDisablesFeishu(t *testing.T) {
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
	cfg.Feishu.Alert.ActionToken = "card-action-token"
	raw, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal config: %v", err)
	}
	repo.values[SettingKeyOpsEmailNotificationConfig] = string(raw)

	opsSvc := &OpsService{settingRepo: repo}
	evaluator := NewOpsAlertEvaluatorService(opsSvc, &opsRepoMock{}, nil, nil, nil)
	rule := &OpsAlertRule{ID: 1, Name: "High error rate", Severity: "critical", MetricType: "error_rate", Operator: ">", Threshold: 5, NotifyEmail: false, NotifyFeishu: false}
	value := 9.2
	event := &OpsAlertEvent{ID: 10, RuleID: 1, Status: "firing", MetricValue: &value, FiredAt: time.Now().UTC(), Description: "error_rate above threshold"}

	if sent := evaluator.maybeSendAlertEmail(context.Background(), nil, rule, event); sent {
		t.Fatalf("maybeSendAlertEmail() = true, want false")
	}
	if webhookCalls != 0 {
		t.Fatalf("webhook calls = %d, want 0", webhookCalls)
	}
}

func TestHandleFeishuAlertCardAction_ManualResolvesEvent(t *testing.T) {
	repo := newRuntimeSettingRepoStub()
	cfg := defaultOpsEmailNotificationConfig()
	cfg.Feishu.Alert.ActionToken = "card-action-token"
	raw, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal config: %v", err)
	}
	repo.values[SettingKeyOpsEmailNotificationConfig] = string(raw)

	var gotEventID int64
	var gotStatus string
	var gotResolvedAt bool
	opsSvc := &OpsService{
		settingRepo: repo,
		opsRepo: &opsRepoMock{
			UpdateAlertEventStatusFn: func(ctx context.Context, eventID int64, status string, resolvedAt *time.Time) error {
				gotEventID = eventID
				gotStatus = status
				gotResolvedAt = resolvedAt != nil
				return nil
			},
		},
	}

	result, err := opsSvc.HandleFeishuAlertCardAction(context.Background(), OpsFeishuAlertCardAction{
		Action:  "resolve_alert",
		EventID: 42,
		Token:   "card-action-token",
	})
	if err != nil {
		t.Fatalf("HandleFeishuAlertCardAction() error = %v", err)
	}
	if result == nil || result.Status != "manual_resolved" {
		t.Fatalf("result = %#v, want manual_resolved", result)
	}
	if gotEventID != 42 || gotStatus != OpsAlertStatusManualResolved || !gotResolvedAt {
		t.Fatalf("update status got event=%d status=%q resolvedAt=%v", gotEventID, gotStatus, gotResolvedAt)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
