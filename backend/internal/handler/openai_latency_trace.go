package handler

import (
	"time"

	"github.com/Wei-Shaw/sub2api/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type openAIChatCompletionsLatencyTraceInput struct {
	RequestStart     time.Time
	Endpoint         string
	Model            string
	Stream           bool
	Status           int
	UserID           int64
	APIKeyID         int64
	GroupID          *int64
	Account          *service.Account
	SwitchCount      int
	RequestBodyBytes int
	Result           *service.OpenAIForwardResult
}

func logOpenAIChatCompletionsLatencyTrace(c *gin.Context, reqLog *zap.Logger, in openAIChatCompletionsLatencyTraceInput) {
	if reqLog == nil {
		return
	}
	if !service.ShouldSampleOpenAIGatewayLatencyTrace() {
		return
	}
	reqLog.Info("openai_gateway_latency", buildOpenAIChatCompletionsLatencyFields(c, in, time.Now())...)
}

func buildOpenAIChatCompletionsLatencyFields(c *gin.Context, in openAIChatCompletionsLatencyTraceInput, now time.Time) []zap.Field {
	endpoint := in.Endpoint
	if endpoint == "" && c != nil {
		endpoint = c.FullPath()
		if endpoint == "" && c.Request != nil && c.Request.URL != nil {
			endpoint = c.Request.URL.Path
		}
	}

	fields := []zap.Field{
		zap.String("event", "openai_gateway_latency"),
		zap.String("endpoint", endpoint),
		zap.String("model", in.Model),
		zap.Bool("stream", in.Stream),
		zap.Int("status", in.Status),
		zap.Int64("user_id", in.UserID),
		zap.Int64("api_key_id", in.APIKeyID),
		zap.Int("switch_count", in.SwitchCount),
		zap.Int("request_body_bytes", in.RequestBodyBytes),
	}
	if !in.RequestStart.IsZero() {
		fields = append(fields, zap.Int64("total_ms", now.Sub(in.RequestStart).Milliseconds()))
	}
	if in.GroupID != nil {
		fields = append(fields, zap.Int64("group_id", *in.GroupID))
	}
	if in.Account != nil {
		fields = append(fields,
			zap.Int64("account_id", in.Account.ID),
			zap.String("account_platform", in.Account.Platform),
			zap.String("account_type", in.Account.Type),
		)
	}
	fields = appendOpenAIChatCompletionsLatencyField(fields, c, service.OpsAuthLatencyMsKey, "auth_ms")
	fields = appendOpenAIChatCompletionsLatencyField(fields, c, service.OpsRoutingLatencyMsKey, "routing_ms")
	fields = appendOpenAIChatCompletionsLatencyField(fields, c, service.OpsUpstreamLatencyMsKey, "upstream_ms")
	fields = appendOpenAIChatCompletionsLatencyField(fields, c, service.OpsResponseLatencyMsKey, "response_ms")
	fields = appendOpenAIChatCompletionsLatencyField(fields, c, service.OpsTimeToFirstTokenMsKey, "first_token_ms")
	fields = appendOpenAIChatCompletionsLatencyField(fields, c, service.OpsOpenAIWSQueueWaitMsKey, "openai_ws_queue_wait_ms")
	fields = appendOpenAIChatCompletionsLatencyField(fields, c, service.OpsOpenAIWSConnPickMsKey, "openai_ws_conn_pick_ms")

	if c != nil {
		if v, ok := c.Get(service.OpsOpenAIWSConnReusedKey); ok {
			if reused, ok := v.(bool); ok {
				fields = append(fields, zap.Bool("openai_ws_conn_reused", reused))
			}
		}
		if v, ok := c.Get(service.OpsOpenAIWSConnIDKey); ok {
			if connID, ok := v.(string); ok && connID != "" {
				fields = append(fields, zap.String("openai_ws_conn_id", connID))
			}
		}
	}

	if in.Result != nil {
		if in.Result.RequestID != "" {
			fields = append(fields, zap.String("upstream_request_id", in.Result.RequestID))
		}
		if in.Result.ResponseID != "" {
			fields = append(fields, zap.String("upstream_response_id", in.Result.ResponseID))
		}
		if in.Result.UpstreamModel != "" {
			fields = append(fields, zap.String("upstream_model", in.Result.UpstreamModel))
		}
		fields = append(fields, zap.Bool("openai_ws_mode", in.Result.OpenAIWSMode))
	}

	return fields
}

func appendOpenAIChatCompletionsLatencyField(fields []zap.Field, c *gin.Context, contextKey, fieldName string) []zap.Field {
	if v, ok := getContextInt64(c, contextKey); ok {
		fields = append(fields, zap.Int64(fieldName, v))
	}
	return fields
}
