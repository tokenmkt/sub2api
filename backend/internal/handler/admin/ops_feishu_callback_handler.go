package admin

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Wei-Shaw/sub2api/internal/pkg/response"
	"github.com/Wei-Shaw/sub2api/internal/service"
	"github.com/gin-gonic/gin"
)

// HandleFeishuAlertCallback handles Feishu interactive card callbacks.
// POST /api/v1/admin/ops/feishu/alert-callback
func (h *OpsHandler) HandleFeishuAlertCallback(c *gin.Context) {
	if h.opsService == nil {
		response.Error(c, http.StatusServiceUnavailable, "Ops service not available")
		return
	}

	var payload map[string]json.RawMessage
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.BadRequest(c, "Invalid request body")
		return
	}

	if challenge, ok := readJSONString(payload["challenge"]); ok && challenge != "" {
		c.JSON(http.StatusOK, gin.H{"challenge": challenge})
		return
	}

	action, err := parseFeishuAlertCardAction(payload)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.opsService.HandleFeishuAlertCardAction(c.Request.Context(), action)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"toast": gin.H{
			"type":    "success",
			"content": "告警已标记为已处理",
		},
		"data": result,
	})
}

func parseFeishuAlertCardAction(payload map[string]json.RawMessage) (service.OpsFeishuAlertCardAction, error) {
	var action service.OpsFeishuAlertCardAction

	if rawAction, ok := payload["action"]; ok {
		var actionPayload struct {
			Value service.OpsFeishuAlertCardAction `json:"value"`
		}
		if err := json.Unmarshal(rawAction, &actionPayload); err == nil {
			action = actionPayload.Value
		}
	}

	if action.Action == "" {
		if rawValue, ok := payload["value"]; ok {
			_ = json.Unmarshal(rawValue, &action)
		}
	}

	action.Action = strings.TrimSpace(action.Action)
	action.Token = strings.TrimSpace(action.Token)
	if action.Action == "" || action.EventID <= 0 || action.Token == "" {
		return action, errInvalidFeishuAlertAction
	}
	return action, nil
}

var errInvalidFeishuAlertAction = &feishuAlertCallbackError{message: "Invalid Feishu alert action"}

type feishuAlertCallbackError struct {
	message string
}

func (e *feishuAlertCallbackError) Error() string {
	return e.message
}

func readJSONString(raw json.RawMessage) (string, bool) {
	if len(raw) == 0 {
		return "", false
	}
	var value string
	if err := json.Unmarshal(raw, &value); err != nil {
		return "", false
	}
	return strings.TrimSpace(value), true
}
