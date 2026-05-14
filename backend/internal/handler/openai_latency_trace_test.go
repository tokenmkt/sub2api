package handler

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Wei-Shaw/sub2api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestBuildOpenAIChatCompletionsLatencyFieldsIncludesStageTimings(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("POST", "/v1/chat/completions", nil)
	c.Params = gin.Params{{Key: "path", Value: "/v1/chat/completions"}}
	c.Set(service.OpsAuthLatencyMsKey, int64(11))
	c.Set(service.OpsRoutingLatencyMsKey, int64(22))
	c.Set(service.OpsUpstreamLatencyMsKey, int64(333))
	c.Set(service.OpsResponseLatencyMsKey, int64(44))
	c.Set(service.OpsTimeToFirstTokenMsKey, int64(555))
	c.Set(service.OpsOpenAIWSQueueWaitMsKey, int64(6))
	c.Set(service.OpsOpenAIWSConnPickMsKey, int64(7))
	c.Set(service.OpsOpenAIWSConnReusedKey, true)
	c.Set(service.OpsOpenAIWSConnIDKey, "conn-1")

	groupID := int64(9)
	start := time.Unix(100, 0)
	now := start.Add(1234 * time.Millisecond)
	firstToken := 555
	fields := buildOpenAIChatCompletionsLatencyFields(c, openAIChatCompletionsLatencyTraceInput{
		RequestStart:     start,
		Endpoint:         "/v1/chat/completions",
		Model:            "gpt-5.5",
		Stream:           true,
		Status:           200,
		UserID:           7,
		APIKeyID:         8,
		GroupID:          &groupID,
		Account:          &service.Account{ID: 42, Platform: service.PlatformOpenAI, Type: service.AccountTypeAPIKey},
		SwitchCount:      2,
		RequestBodyBytes: 321,
		Result: &service.OpenAIForwardResult{
			RequestID:     "req-upstream",
			ResponseID:    "resp-upstream",
			UpstreamModel: "gpt-5.5",
			Stream:        true,
			OpenAIWSMode:  true,
			FirstTokenMs:  &firstToken,
		},
	}, now)

	got := zapFieldsToMap(fields)
	require.Equal(t, "openai_gateway_latency", got["event"])
	require.Equal(t, "/v1/chat/completions", got["endpoint"])
	require.Equal(t, "gpt-5.5", got["model"])
	require.Equal(t, true, got["stream"])
	require.Equal(t, int64(1234), got["total_ms"])
	require.Equal(t, int64(11), got["auth_ms"])
	require.Equal(t, int64(22), got["routing_ms"])
	require.Equal(t, int64(333), got["upstream_ms"])
	require.Equal(t, int64(44), got["response_ms"])
	require.Equal(t, int64(555), got["first_token_ms"])
	require.Equal(t, int64(6), got["openai_ws_queue_wait_ms"])
	require.Equal(t, int64(7), got["openai_ws_conn_pick_ms"])
	require.Equal(t, true, got["openai_ws_conn_reused"])
	require.Equal(t, "conn-1", got["openai_ws_conn_id"])
	require.Equal(t, int64(42), got["account_id"])
	require.Equal(t, "openai", got["account_platform"])
	require.Equal(t, "apikey", got["account_type"])
	require.Equal(t, int64(2), got["switch_count"])
	require.Equal(t, int64(321), got["request_body_bytes"])
	require.Equal(t, "req-upstream", got["upstream_request_id"])
	require.Equal(t, "resp-upstream", got["upstream_response_id"])
	require.Equal(t, "gpt-5.5", got["upstream_model"])
	require.Equal(t, true, got["openai_ws_mode"])
}

func zapFieldsToMap(fields []zap.Field) map[string]any {
	out := make(map[string]any, len(fields))
	for _, field := range fields {
		switch field.Type {
		case zapcore.StringType:
			out[field.Key] = field.String
		case zapcore.Int64Type:
			out[field.Key] = field.Integer
		case zapcore.Int32Type:
			out[field.Key] = int32(field.Integer)
		case zapcore.Int16Type:
			out[field.Key] = int16(field.Integer)
		case zapcore.Int8Type:
			out[field.Key] = int8(field.Integer)
		case zapcore.BoolType:
			out[field.Key] = field.Integer == 1
		}
	}
	return out
}
