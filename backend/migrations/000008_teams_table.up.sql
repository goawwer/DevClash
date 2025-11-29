CREATE TABLE teams (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    creator_id UUID REFERENCES users (id) ON DELETE SET NULL,
    leader_id UUID REFERENCES users (id) ON DELETE SET NULL,
    name VARCHAR(80) NOT NULL UNIQUE,
    team_status VARCHAR(150),
    description TEXT,
    team_picture_url TEXT,
    participations_count INT NOT NULL DEFAULT 0,
    wins_count INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    disabled BOOLEAN DEFAULT FALSE
);