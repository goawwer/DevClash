CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    organizer_id UUID REFERENCES organizers (id) ON DELETE CASCADE,
    type_id UUID REFERENCES event_types (id) NOT NULL,
    title VARCHAR(150) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp,
    is_finished BOOLEAN DEFAULT FALSE
);