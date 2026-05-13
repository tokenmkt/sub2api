//go:build unit

package service

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmailService_GetEmailConfig_Resend(t *testing.T) {
	repo := &settingRepoStub{values: map[string]string{
		SettingKeyEmailProvider: "resend",
		SettingKeyResendAPIKey:  "  re_123  ",
		SettingKeySMTPFrom:      " no-reply@example.com ",
		SettingKeySMTPFromName:  " Example App ",
	}}
	svc := NewEmailService(repo, nil)

	cfg, err := svc.GetEmailConfig(context.Background())

	require.NoError(t, err)
	require.Equal(t, EmailProviderResend, cfg.Provider)
	require.Equal(t, "re_123", cfg.Resend.APIKey)
	require.Equal(t, "no-reply@example.com", cfg.Resend.From)
	require.Equal(t, "Example App", cfg.Resend.FromName)
}

func TestEmailService_SendEmailWithConfig_ResendPostsHTTPSPayload(t *testing.T) {
	var gotAuth string
	var gotPayload map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/emails", r.URL.Path)
		gotAuth = r.Header.Get("Authorization")
		require.Equal(t, "application/json", r.Header.Get("Content-Type"))
		require.NoError(t, json.NewDecoder(r.Body).Decode(&gotPayload))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"email_123"}`))
	}))
	defer server.Close()

	svc := NewEmailService(nil, nil)
	svc.SetResendAPIBaseURLForTest(server.URL)
	err := svc.SendEmailWithConfig(&EmailConfig{
		Provider: EmailProviderResend,
		Resend: ResendConfig{
			APIKey:   "re_secret",
			From:     "no-reply@example.com",
			FromName: "Example App",
		},
	}, "user@example.com", "Hello", "<p>World</p>")

	require.NoError(t, err)
	require.Equal(t, "Bearer re_secret", gotAuth)
	require.Equal(t, "Example App <no-reply@example.com>", gotPayload["from"])
	require.Equal(t, []any{"user@example.com"}, gotPayload["to"])
	require.Equal(t, "Hello", gotPayload["subject"])
	require.Equal(t, "<p>World</p>", gotPayload["html"])
}

func TestEmailService_SendEmailWithConfig_ResendReturnsAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"message":"invalid api key"}`))
	}))
	defer server.Close()

	svc := NewEmailService(nil, nil)
	svc.SetResendAPIBaseURLForTest(server.URL)
	err := svc.SendEmailWithConfig(&EmailConfig{
		Provider: EmailProviderResend,
		Resend: ResendConfig{
			APIKey: "bad",
			From:   "no-reply@example.com",
		},
	}, "user@example.com", "Hello", "<p>World</p>")

	require.Error(t, err)
	require.Contains(t, err.Error(), "resend send failed")
	require.Contains(t, err.Error(), "invalid api key")
}
