CREATE TABLE event_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(150)
);

INSERT INTO
    event_types (name)
VALUES ('hackathon'),
    ('exhibition'),
    ('camp'),
    ('conference');