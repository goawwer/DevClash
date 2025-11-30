CREATE TABLE event_properties (
    event_id UUID PRIMARY KEY REFERENCES events (id) ON DELETE CASCADE,
    is_online BOOLEAN NOT NULL DEFAULT TRUE,
    is_free BOOLEAN NOT NULL DEFAULT TRUE,
    number_of_teams INT NOT NULL DEFAULT 4,
    team_size INT NOT NULL DEFAULT 4
);