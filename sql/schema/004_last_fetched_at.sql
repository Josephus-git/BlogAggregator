-- +gooseUp
ALTER TABLE feeds
ADD COLUMN last_fetched_at TIMESTAMP;

-- +gooseDown
ALTER TABLE feeds
DROP COLUMN last_fetched_at;