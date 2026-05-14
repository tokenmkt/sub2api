package service

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	OpenAIGatewayLatencyTraceEnabledEnv    = "OPENAI_GATEWAY_LATENCY_TRACE_ENABLED"
	OpenAIGatewayLatencyTraceSampleRateEnv = "OPENAI_GATEWAY_LATENCY_TRACE_SAMPLE_RATE"
	UpstreamHTTPTraceEnabledEnv            = "UPSTREAM_HTTP_TRACE_ENABLED"
	UpstreamHTTPTraceSampleRateEnv         = "UPSTREAM_HTTP_TRACE_SAMPLE_RATE"
)

func ShouldSampleOpenAIGatewayLatencyTrace() bool {
	return shouldSampleTrace(OpenAIGatewayLatencyTraceEnabledEnv, OpenAIGatewayLatencyTraceSampleRateEnv)
}

func ShouldSampleUpstreamHTTPTrace() bool {
	return shouldSampleTrace(UpstreamHTTPTraceEnabledEnv, UpstreamHTTPTraceSampleRateEnv)
}

func shouldSampleTrace(enabledEnv, sampleRateEnv string) bool {
	if isTraceExplicitlyDisabled(os.Getenv(enabledEnv)) {
		return false
	}
	rate := parseTraceSampleRate(os.Getenv(sampleRateEnv))
	if rate <= 0 {
		return false
	}
	if rate >= 1 {
		return true
	}
	return rand.Float64() < rate
}

func isTraceExplicitlyDisabled(raw string) bool {
	switch strings.ToLower(strings.TrimSpace(raw)) {
	case "0", "false", "off", "no", "disabled":
		return true
	default:
		return false
	}
}

func parseTraceSampleRate(raw string) float64 {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return 1
	}
	rate, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 1
	}
	if rate < 0 {
		return 0
	}
	if rate > 1 {
		return 1
	}
	return rate
}
