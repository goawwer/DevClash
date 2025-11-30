CREATE TABLE events_teams (
    event_id UUID REFERENCES events (id) ON DELETE CASCADE,
    team_id UUID REFERENCES teams (id) ON DELETE CASCADE,
    joined_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (event_id, team_id)
);