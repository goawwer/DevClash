CREATE TABLE teams_members (
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    team_id UUID REFERENCES teams (id) ON DELETE CASCADE,
    role UUID REFERENCES teams_roles (id),
    joined_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (user_id, team_id)
);