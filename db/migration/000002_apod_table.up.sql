CREATE TABLE IF NOT EXISTS day_info.apod (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP,
    explanation TEXT,
    media_type VARCHAR(255),
    service_version VARCHAR(255),
    title VARCHAR(255),
    image BYTEA,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);