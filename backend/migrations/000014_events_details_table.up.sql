CREATE TABLE events_details (
    event_id UUID PRIMARY KEY REFERENCES events (id) ON DELETE CASCADE,
    event_picture_url TEXT NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ CHECK (end_time > start_time) NOT NULL,
    description TEXT NOT NULL,
    prize VARCHAR(200)
);