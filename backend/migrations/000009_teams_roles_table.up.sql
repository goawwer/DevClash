CREATE TABLE teams_roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(150) NOT NULL UNIQUE
);

INSERT INTO
    teams_roles (name)
VALUES ('Frontend'),
    ('Backend'),
    ('Testing'),
    ('Lead'),
    ('DBA'),
    ('SRE'),
    ('DevOps');