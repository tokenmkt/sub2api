package service

import (
	"context"
	"errors"
	"strings"
	"time"
)

const OpsFeishuAlertActionResolve = "resolve_alert"

type OpsFeishuAlertCardAction struct {
	Action  string `json:"action"`
	EventID int64  `json:"event_id"`
	RuleID  int64  `json:"rule_id,omitempty"`
	Token   string `json:"token"`
}

type OpsFeishuAlertCardActionResult struct {
	Handled bool   `json:"handled"`
	Status  string `json:"status,omitempty"`
}

func (s *OpsService) HandleFeishuAlertCardAction(ctx context.Context, action OpsFeishuAlertCardAction) (*OpsFeishuAlertCardActionResult, error) {
	if s == nil || s.opsRepo == nil {
		return nil, errors.New("ops repository not initialized")
	}
	if ctx == nil {
		ctx = context.Background()
	}

	if strings.TrimSpace(action.Action) != OpsFeishuAlertActionResolve {
		return nil, errors.New("unsupported feishu alert action")
	}
	if action.EventID <= 0 {
		return nil, errors.New("invalid alert event id")
	}

	cfg, err := s.GetEmailNotificationConfig(ctx)
	if err != nil {
		return nil, err
	}
	expectedToken := strings.TrimSpace(cfg.Feishu.Alert.ActionToken)
	if expectedToken == "" || strings.TrimSpace(action.Token) != expectedToken {
		return nil, errors.New("invalid feishu alert action token")
	}

	now := time.Now().UTC()
	if err := s.opsRepo.UpdateAlertEventStatus(ctx, action.EventID, OpsAlertStatusManualResolved, &now); err != nil {
		return nil, err
	}
	return &OpsFeishuAlertCardActionResult{
		Handled: true,
		Status:  OpsAlertStatusManualResolved,
	}, nil
}
