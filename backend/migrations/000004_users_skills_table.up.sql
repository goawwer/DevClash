CREATE TABLE users_skills (
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    technology_id UUID REFERENCES technologies (id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, technology_id)
);