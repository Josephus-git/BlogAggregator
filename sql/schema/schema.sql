ALTER TABLE feeds
ADD COLUMN last_fetched_at TIMESTAMP;

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT NULL,
    published_at TIMESTAMP NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);