package service

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"strings"
	"sync"
	"time"

	"github.com/Wei-Shaw/sub2api/internal/pkg/logger"
	"go.uber.org/zap"
)

const upstreamHTTPTraceEvent = "upstream_http_trace"

type UpstreamHTTPTraceOptions struct {
	AccountID  int64
	ProxyURL   string
	TLSProfile string
	Now        func() time.Time
}

type UpstreamHTTPTraceRecorder struct {
	mu   sync.Mutex
	now  func() time.Time
	opts UpstreamHTTPTraceOptions

	method string
	scheme string
	host   string
	path   string

	startTime             time.Time
	dnsStart              time.Time
	dnsDone               time.Time
	connectStart          time.Time
	connectDone           time.Time
	tlsStart              time.Time
	tlsDone               time.Time
	gotConn               time.Time
	wroteRequest          time.Time
	firstResponseByte     time.Time
	connectionReused      bool
	connectionWasIdle     bool
	connectionIdle        time.Duration
	remoteAddr            string
	dnsErr                string
	connectErr            string
	tlsErr                string
	wroteRequestErr       string
	firstResponseObserved bool
}

func NewUpstreamHTTPTraceRecorder(req *http.Request, opts UpstreamHTTPTraceOptions) *UpstreamHTTPTraceRecorder {
	if !ShouldSampleUpstreamHTTPTrace() {
		return nil
	}
	now := opts.Now
	if now == nil {
		now = time.Now
	}
	rec := &UpstreamHTTPTraceRecorder{
		now:       now,
		opts:      opts,
		startTime: now(),
	}
	if req != nil {
		rec.method = req.Method
		if req.URL != nil {
			rec.scheme = req.URL.Scheme
			rec.host = req.URL.Hostname()
			rec.path = req.URL.EscapedPath()
			if rec.path == "" {
				rec.path = "/"
			}
		}
	}
	return rec
}

func (r *UpstreamHTTPTraceRecorder) WithClientTrace(req *http.Request) *http.Request {
	if r == nil || req == nil {
		return req
	}
	trace := &httptrace.ClientTrace{
		DNSStart: func(_ httptrace.DNSStartInfo) {
			r.setTime(&r.dnsStart)
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			r.mu.Lock()
			r.dnsDone = r.now()
			if info.Err != nil {
				r.dnsErr = info.Err.Error()
			}
			r.mu.Unlock()
		},
		ConnectStart: func(_, _ string) {
			r.setTime(&r.connectStart)
		},
		ConnectDone: func(_, _ string, err error) {
			r.mu.Lock()
			r.connectDone = r.now()
			if err != nil {
				r.connectErr = err.Error()
			}
			r.mu.Unlock()
		},
		TLSHandshakeStart: func() {
			r.setTime(&r.tlsStart)
		},
		TLSHandshakeDone: func(_ tls.ConnectionState, err error) {
			r.mu.Lock()
			r.tlsDone = r.now()
			if err != nil {
				r.tlsErr = err.Error()
			}
			r.mu.Unlock()
		},
		GotConn: func(info httptrace.GotConnInfo) {
			r.mu.Lock()
			r.gotConn = r.now()
			r.connectionReused = info.Reused
			r.connectionWasIdle = info.WasIdle
			r.connectionIdle = info.IdleTime
			if info.Conn != nil && info.Conn.RemoteAddr() != nil {
				r.remoteAddr = info.Conn.RemoteAddr().String()
			}
			r.mu.Unlock()
		},
		WroteRequest: func(info httptrace.WroteRequestInfo) {
			r.mu.Lock()
			r.wroteRequest = r.now()
			if info.Err != nil {
				r.wroteRequestErr = info.Err.Error()
			}
			r.mu.Unlock()
		},
		GotFirstResponseByte: func() {
			r.mu.Lock()
			r.firstResponseByte = r.now()
			r.firstResponseObserved = true
			r.mu.Unlock()
		},
	}
	return req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
}

func (r *UpstreamHTTPTraceRecorder) Log(ctx context.Context, status int, err error, end time.Time) {
	if r == nil {
		return
	}
	logger.FromContext(ctx).Info(upstreamHTTPTraceEvent, r.Fields(status, err, end)...)
}

func (r *UpstreamHTTPTraceRecorder) Fields(status int, err error, end time.Time) []zap.Field {
	if r == nil {
		return nil
	}
	if end.IsZero() {
		end = r.now()
	}
	r.mu.Lock()
	snap := *r
	r.mu.Unlock()

	fields := []zap.Field{
		zap.String("event", upstreamHTTPTraceEvent),
		zap.Int64("account_id", snap.opts.AccountID),
		zap.String("method", snap.method),
		zap.String("scheme", snap.scheme),
		zap.String("host", snap.host),
		zap.String("path", snap.path),
		zap.Int("status", status),
		zap.Int64("total_ms", durationMs(snap.startTime, end)),
		zap.Bool("connection_reused", snap.connectionReused),
		zap.Bool("connection_was_idle", snap.connectionWasIdle),
	}
	if snap.opts.ProxyURL != "" {
		fields = append(fields, zap.Bool("proxy_enabled", true))
	} else {
		fields = append(fields, zap.Bool("proxy_enabled", false))
	}
	if snap.opts.TLSProfile != "" {
		fields = append(fields, zap.String("tls_profile", snap.opts.TLSProfile))
	}
	if snap.connectionIdle > 0 {
		fields = append(fields, zap.Int64("connection_idle_ms", snap.connectionIdle.Milliseconds()))
	}
	if snap.remoteAddr != "" {
		fields = append(fields, zap.String("remote_addr", snap.remoteAddr))
	}
	if err != nil {
		fields = append(fields, zap.String("error", sanitizeTraceError(err.Error())))
	}
	if snap.dnsErr != "" {
		fields = append(fields, zap.String("dns_error", sanitizeTraceError(snap.dnsErr)))
	}
	if snap.connectErr != "" {
		fields = append(fields, zap.String("connect_error", sanitizeTraceError(snap.connectErr)))
	}
	if snap.tlsErr != "" {
		fields = append(fields, zap.String("tls_error", sanitizeTraceError(snap.tlsErr)))
	}
	if snap.wroteRequestErr != "" {
		fields = append(fields, zap.String("write_error", sanitizeTraceError(snap.wroteRequestErr)))
	}
	fields = appendDurationField(fields, "dns_ms", snap.dnsStart, snap.dnsDone)
	fields = appendDurationField(fields, "connect_ms", snap.connectStart, snap.connectDone)
	fields = appendDurationField(fields, "tls_ms", snap.tlsStart, snap.tlsDone)
	if !snap.gotConn.IsZero() {
		fields = append(fields, zap.Int64("wait_conn_ms", durationMs(snap.startTime, snap.gotConn)))
	}
	if !snap.wroteRequest.IsZero() {
		fields = append(fields, zap.Int64("request_write_ms", durationMs(snap.startTime, snap.wroteRequest)))
	}
	if snap.firstResponseObserved {
		fields = append(fields,
			zap.Int64("ttfb_ms", durationMs(snap.startTime, snap.firstResponseByte)),
			zap.Int64("body_read_ms", durationMs(snap.firstResponseByte, end)),
		)
	}
	fields = append(fields, zap.Int64("body_total_ms", durationMs(snap.startTime, end)))
	return fields
}

func (r *UpstreamHTTPTraceRecorder) setTime(target *time.Time) {
	r.mu.Lock()
	*target = r.now()
	r.mu.Unlock()
}

func appendDurationField(fields []zap.Field, name string, start, end time.Time) []zap.Field {
	if start.IsZero() || end.IsZero() {
		return fields
	}
	return append(fields, zap.Int64(name, durationMs(start, end)))
}

func durationMs(start, end time.Time) int64 {
	if start.IsZero() || end.IsZero() || end.Before(start) {
		return 0
	}
	return end.Sub(start).Milliseconds()
}

func sanitizeTraceError(value string) string {
	value = strings.TrimSpace(value)
	if len(value) > 512 {
		return value[:512]
	}
	return value
}
