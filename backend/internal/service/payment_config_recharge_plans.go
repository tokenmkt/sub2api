package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	dbent "github.com/Wei-Shaw/sub2api/ent"
	infraerrors "github.com/Wei-Shaw/sub2api/internal/pkg/errors"
)

type RechargePlan struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	CreditAmount  float64   `json:"credit_amount"`
	OriginalPrice *float64  `json:"original_price,omitempty"`
	Features      string    `json:"features"`
	PurchaseURL   string    `json:"purchase_url"`
	Badge         string    `json:"badge"`
	ForSale       bool      `json:"for_sale"`
	SortOrder     int       `json:"sort_order"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateRechargePlanRequest struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Price         float64  `json:"price"`
	CreditAmount  float64  `json:"credit_amount"`
	OriginalPrice *float64 `json:"original_price"`
	Features      string   `json:"features"`
	PurchaseURL   string   `json:"purchase_url"`
	Badge         string   `json:"badge"`
	ForSale       bool     `json:"for_sale"`
	SortOrder     int      `json:"sort_order"`
}

type UpdateRechargePlanRequest struct {
	Name          *string  `json:"name"`
	Description   *string  `json:"description"`
	Price         *float64 `json:"price"`
	CreditAmount  *float64 `json:"credit_amount"`
	OriginalPrice *float64 `json:"original_price"`
	Features      *string  `json:"features"`
	PurchaseURL   *string  `json:"purchase_url"`
	Badge         *string  `json:"badge"`
	ForSale       *bool    `json:"for_sale"`
	SortOrder     *int     `json:"sort_order"`
}

func (s *PaymentConfigService) ListRechargePlans(ctx context.Context) ([]RechargePlan, error) {
	return s.listRechargePlans(ctx, false)
}

func (s *PaymentConfigService) ListRechargePlansForSale(ctx context.Context) ([]RechargePlan, error) {
	return s.listRechargePlans(ctx, true)
}

func (s *PaymentConfigService) CreateRechargePlan(ctx context.Context, req CreateRechargePlanRequest) (*RechargePlan, error) {
	if err := validateRechargePlanRequired(req.Name, req.Price, req.CreditAmount, req.PurchaseURL, req.OriginalPrice); err != nil {
		return nil, err
	}
	return querySingleRechargePlan(ctx, s.entClient, `
INSERT INTO balance_recharge_plans (
	name, description, price, credit_amount, original_price, features,
	purchase_url, badge, for_sale, sort_order, created_at, updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, name, description, price, credit_amount, original_price, features,
	purchase_url, badge, for_sale, sort_order, created_at, updated_at`,
		strings.TrimSpace(req.Name), req.Description, req.Price, req.CreditAmount,
		nullableFloat(req.OriginalPrice), req.Features, strings.TrimSpace(req.PurchaseURL),
		req.Badge, req.ForSale, req.SortOrder,
	)
}

func (s *PaymentConfigService) UpdateRechargePlan(ctx context.Context, id int64, req UpdateRechargePlanRequest) (*RechargePlan, error) {
	if id <= 0 {
		return nil, infraerrors.BadRequest("INVALID_ID", "invalid recharge plan id")
	}
	existing, err := s.GetRechargePlan(ctx, id)
	if err != nil {
		return nil, err
	}
	merged := CreateRechargePlanRequest{
		Name:          existing.Name,
		Description:   existing.Description,
		Price:         existing.Price,
		CreditAmount:  existing.CreditAmount,
		OriginalPrice: existing.OriginalPrice,
		Features:      existing.Features,
		PurchaseURL:   existing.PurchaseURL,
		Badge:         existing.Badge,
		ForSale:       existing.ForSale,
		SortOrder:     existing.SortOrder,
	}
	if req.Name != nil {
		merged.Name = *req.Name
	}
	if req.Description != nil {
		merged.Description = *req.Description
	}
	if req.Price != nil {
		merged.Price = *req.Price
	}
	if req.CreditAmount != nil {
		merged.CreditAmount = *req.CreditAmount
	}
	if req.OriginalPrice != nil {
		merged.OriginalPrice = req.OriginalPrice
	}
	if req.Features != nil {
		merged.Features = *req.Features
	}
	if req.PurchaseURL != nil {
		merged.PurchaseURL = *req.PurchaseURL
	}
	if req.Badge != nil {
		merged.Badge = *req.Badge
	}
	if req.ForSale != nil {
		merged.ForSale = *req.ForSale
	}
	if req.SortOrder != nil {
		merged.SortOrder = *req.SortOrder
	}
	if err := validateRechargePlanRequired(merged.Name, merged.Price, merged.CreditAmount, merged.PurchaseURL, merged.OriginalPrice); err != nil {
		return nil, err
	}

	return querySingleRechargePlan(ctx, s.entClient, `
UPDATE balance_recharge_plans
SET name = $2, description = $3, price = $4, credit_amount = $5, original_price = $6,
	features = $7, purchase_url = $8, badge = $9, for_sale = $10, sort_order = $11, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, name, description, price, credit_amount, original_price, features,
	purchase_url, badge, for_sale, sort_order, created_at, updated_at`,
		id, strings.TrimSpace(merged.Name), merged.Description, merged.Price,
		merged.CreditAmount, nullableFloat(merged.OriginalPrice), merged.Features,
		strings.TrimSpace(merged.PurchaseURL), merged.Badge, merged.ForSale, merged.SortOrder,
	)
}

func (s *PaymentConfigService) DeleteRechargePlan(ctx context.Context, id int64) error {
	if id <= 0 {
		return infraerrors.BadRequest("INVALID_ID", "invalid recharge plan id")
	}
	res, err := s.entClient.ExecContext(ctx, `DELETE FROM balance_recharge_plans WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete recharge plan: %w", err)
	}
	affected, err := res.RowsAffected()
	if err == nil && affected == 0 {
		return infraerrors.NotFound("RECHARGE_PLAN_NOT_FOUND", "recharge plan not found")
	}
	return nil
}

func (s *PaymentConfigService) GetRechargePlan(ctx context.Context, id int64) (*RechargePlan, error) {
	plan, err := querySingleRechargePlan(ctx, s.entClient, `
SELECT id, name, description, price, credit_amount, original_price, features,
	purchase_url, badge, for_sale, sort_order, created_at, updated_at
FROM balance_recharge_plans
WHERE id = $1`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, infraerrors.NotFound("RECHARGE_PLAN_NOT_FOUND", "recharge plan not found")
	}
	return plan, err
}

func (s *PaymentConfigService) listRechargePlans(ctx context.Context, onlyForSale bool) ([]RechargePlan, error) {
	query := `
SELECT id, name, description, price, credit_amount, original_price, features,
	purchase_url, badge, for_sale, sort_order, created_at, updated_at
FROM balance_recharge_plans`
	if onlyForSale {
		query += ` WHERE for_sale = true`
	}
	query += ` ORDER BY sort_order ASC, id ASC`
	rows, err := s.entClient.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("list recharge plans: %w", err)
	}
	defer func() { _ = rows.Close() }()

	plans := []RechargePlan{}
	for rows.Next() {
		plan, err := scanRechargePlan(rows)
		if err != nil {
			return nil, err
		}
		plans = append(plans, *plan)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return plans, nil
}

type rechargePlanScanner interface {
	Scan(dest ...any) error
}

func scanRechargePlan(scanner rechargePlanScanner) (*RechargePlan, error) {
	var plan RechargePlan
	var original sql.NullFloat64
	var createdAt any
	var updatedAt any
	if err := scanner.Scan(
		&plan.ID, &plan.Name, &plan.Description, &plan.Price, &plan.CreditAmount,
		&original, &plan.Features, &plan.PurchaseURL, &plan.Badge, &plan.ForSale,
		&plan.SortOrder, &createdAt, &updatedAt,
	); err != nil {
		return nil, err
	}
	if original.Valid {
		plan.OriginalPrice = &original.Float64
	}
	plan.CreatedAt = parseDBTime(createdAt)
	plan.UpdatedAt = parseDBTime(updatedAt)
	return &plan, nil
}

func querySingleRechargePlan(ctx context.Context, client *dbent.Client, query string, args ...any) (*RechargePlan, error) {
	rows, err := client.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}
		return nil, sql.ErrNoRows
	}
	return scanRechargePlan(rows)
}

func parseDBTime(value any) time.Time {
	switch v := value.(type) {
	case time.Time:
		return v
	case string:
		for _, layout := range []string{time.RFC3339Nano, "2006-01-02 15:04:05.999999999-07:00", "2006-01-02 15:04:05"} {
			if parsed, err := time.Parse(layout, v); err == nil {
				return parsed
			}
		}
	case []byte:
		return parseDBTime(string(v))
	}
	return time.Time{}
}

func validateRechargePlanRequired(name string, price, creditAmount float64, purchaseURL string, originalPrice *float64) error {
	if strings.TrimSpace(name) == "" {
		return infraerrors.BadRequest("RECHARGE_PLAN_NAME_REQUIRED", "recharge plan name is required")
	}
	if price <= 0 {
		return infraerrors.BadRequest("RECHARGE_PLAN_PRICE_INVALID", "price must be > 0")
	}
	if creditAmount <= 0 {
		return infraerrors.BadRequest("RECHARGE_PLAN_CREDIT_INVALID", "credit amount must be > 0")
	}
	if originalPrice != nil && *originalPrice < 0 {
		return infraerrors.BadRequest("RECHARGE_PLAN_ORIGINAL_PRICE_INVALID", "original price must be >= 0")
	}
	if !isSafePurchaseURL(purchaseURL) {
		return infraerrors.BadRequest("RECHARGE_PLAN_URL_INVALID", "purchase url must be http or https")
	}
	return nil
}

func isSafePurchaseURL(raw string) bool {
	parsed, err := url.ParseRequestURI(strings.TrimSpace(raw))
	if err != nil {
		return false
	}
	return parsed.Scheme == "http" || parsed.Scheme == "https"
}

func nullableFloat(v *float64) any {
	if v == nil {
		return nil
	}
	return *v
}

func ensureRechargePlansTable(ctx context.Context, client *dbent.Client) error {
	_, err := client.ExecContext(ctx, `
CREATE TABLE IF NOT EXISTS balance_recharge_plans (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(100) NOT NULL,
	description TEXT NOT NULL DEFAULT '',
	price DECIMAL(20,2) NOT NULL,
	credit_amount DECIMAL(20,2) NOT NULL,
	original_price DECIMAL(20,2),
	features TEXT NOT NULL DEFAULT '',
	purchase_url TEXT NOT NULL,
	badge VARCHAR(50) NOT NULL DEFAULT '',
	for_sale BOOLEAN NOT NULL DEFAULT true,
	sort_order INTEGER NOT NULL DEFAULT 0,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
)`)
	return err
}
