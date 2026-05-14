-- External balance recharge plans displayed on the user purchase page.
--
-- These plans are not platform payment products. They only provide a managed
-- catalog of external purchase links.

SET LOCAL lock_timeout = '5s';
SET LOCAL statement_timeout = '10min';

CREATE TABLE IF NOT EXISTS balance_recharge_plans (
    id BIGSERIAL PRIMARY KEY,
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
);

CREATE INDEX IF NOT EXISTS idx_balance_recharge_plans_for_sale
    ON balance_recharge_plans (for_sale);

CREATE INDEX IF NOT EXISTS idx_balance_recharge_plans_sort_order
    ON balance_recharge_plans (sort_order, id);
