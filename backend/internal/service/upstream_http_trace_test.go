package service

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptrace"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestUpstreamHTTPTraceFieldsIncludeNetworkPhaseDurations(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/chat/completions?api_key=secret", nil)
	require.NoError(t, err)

	now := time.Unix(100, 0)
	rec := NewUpstreamHTTPTraceRecorder(req, UpstreamHTTPTraceOptions{
		AccountID:  42,
		TLSProfile: "chrome",
		Now: func() time.Time {
			return now
		},
	})
	tracedReq := rec.WithClientTrace(req)
	trace := httptrace.ContextClientTrace(tracedReq.Context())
	require.NotNil(t, trace)

	now = now.Add(1 * time.Millisecond)
	trace.DNSStart(httptrace.DNSStartInfo{})
	now = now.Add(5 * time.Millisecond)
	trace.DNSDone(httptrace.DNSDoneInfo{})
	now = now.Add(1 * time.Millisecond)
	trace.ConnectStart("tcp", "api.openai.com:443")
	now = now.Add(20 * time.Millisecond)
	trace.ConnectDone("tcp", "api.openai.com:443", nil)
	now = now.Add(3 * time.Millisecond)
	trace.TLSHandshakeStart()
	now = now.Add(15 * time.Millisecond)
	trace.TLSHandshakeDone(tls.ConnectionState{}, nil)
	now = now.Add(2 * time.Millisecond)
	trace.GotConn(httptrace.GotConnInfo{
		Reused:   true,
		WasIdle:  true,
		IdleTime: 2 * time.Second,
		Conn:     fakeTraceConn{remote: fakeAddr("198.51.100.10:443")},
	})
	now = now.Add(10 * time.Millisecond)
	trace.WroteRequest(httptrace.WroteRequestInfo{})
	now = now.Add(51 * time.Millisecond)
	trace.GotFirstResponseByte()
	now = now.Add(89 * time.Millisecond)

	got := zapFieldsToMapForUpstreamTrace(rec.Fields(200, nil, now))
	require.Equal(t, "upstream_http_trace", got["event"])
	require.Equal(t, int64(42), got["account_id"])
	require.Equal(t, "POST", got["method"])
	require.Equal(t, "https", got["scheme"])
	require.Equal(t, "api.openai.com", got["host"])
	require.Equal(t, "/v1/chat/completions", got["path"])
	require.Equal(t, true, got["connection_reused"])
	require.Equal(t, true, got["connection_was_idle"])
	require.Equal(t, int64(2000), got["connection_idle_ms"])
	require.Equal(t, "198.51.100.10:443", got["remote_addr"])
	require.Equal(t, "chrome", got["tls_profile"])
	require.Equal(t, int64(197), got["total_ms"])
	require.Equal(t, int64(5), got["dns_ms"])
	require.Equal(t, int64(20), got["connect_ms"])
	require.Equal(t, int64(15), got["tls_ms"])
	require.Equal(t, int64(108), got["ttfb_ms"])
	require.Equal(t, int64(197), got["body_total_ms"])
	require.Equal(t, int64(89), got["body_read_ms"])
	require.NotContains(t, got, "query")
}

func TestUpstreamHTTPTraceSamplingConfigRespectsEnvSwitchAndRate(t *testing.T) {
	t.Setenv("UPSTREAM_HTTP_TRACE_ENABLED", "false")
	t.Setenv("UPSTREAM_HTTP_TRACE_SAMPLE_RATE", "1")
	require.False(t, ShouldSampleUpstreamHTTPTrace())

	t.Setenv("UPSTREAM_HTTP_TRACE_ENABLED", "true")
	t.Setenv("UPSTREAM_HTTP_TRACE_SAMPLE_RATE", "0")
	require.False(t, ShouldSampleUpstreamHTTPTrace())

	t.Setenv("UPSTREAM_HTTP_TRACE_SAMPLE_RATE", "1")
	require.True(t, ShouldSampleUpstreamHTTPTrace())
}

type fakeAddr string

func (a fakeAddr) Network() string { return "tcp" }
func (a fakeAddr) String() string  { return string(a) }

type fakeTraceConn struct {
	net.Conn
	remote net.Addr
}

func (c fakeTraceConn) RemoteAddr() net.Addr { return c.remote }

func zapFieldsToMapForUpstreamTrace(fields []zap.Field) map[string]any {
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
