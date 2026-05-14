-- Ops Monitoring: split alert rule Feishu notifications from email notifications.
--
-- Existing rules keep the previous effective behavior by enabling Feishu
-- notifications by default at the rule level. Global Feishu settings still
-- control whether webhooks are actually sent.

SET LOCAL lock_timeout = '5s';
SET LOCAL statement_timeout = '10min';

ALTER TABLE ops_alert_rules
    ADD COLUMN IF NOT EXISTS notify_feishu BOOLEAN NOT NULL DEFAULT true;
