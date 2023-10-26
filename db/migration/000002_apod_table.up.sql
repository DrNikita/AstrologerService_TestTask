CREATE TABLE IF NOT EXISTS day_info.apod (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    explanation TEXT NOT NULL,
    image TEXT NOT NULL,
    media_type TEXT NOT NULL,
    service_version TEXT NOT NULL
);