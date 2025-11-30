CREATE TABLE events_technologies (
    event_id UUID REFERENCES events (id) ON DELETE CASCADE,
    technology_id UUID REFERENCES technologies (id) ON DELETE CASCADE,
    PRIMARY KEY (event_id, technology_id)
);