CREATE TABLE day_info.apod (
    id SERIAL PRIMARY KEY,
    copyright TEXT,
    date TIMESTAMP,
    explanation TEXT,
    media_type TEXT,
    service_version TEXT,
    title TEXT,
    image BYTEA,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
