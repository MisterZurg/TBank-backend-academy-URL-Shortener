-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE DATABASE IF NOT EXISTS tbank_academy;

CREATE TABLE IF NOT EXISTS tbank_academy.short_to_long (
    short_url String,          -- идентификатор платежа
    long_url String           -- дата платежа
) ENGINE = ReplacingMergeTree()
    ORDER BY short_url; -- в качестве ключа дедупликации


-- CREATE MATERIALIZED VIEW tbank_academy.short_to_long_mv
--     TO tinkoff_academy.short_to_long AS
-- SELECT short_url, long_url
-- FROM tbank_academy.short_to_long;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
