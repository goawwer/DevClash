CREATE TABLE users_skills (
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    tech_stack_id UUID REFERENCES tech_stacks (id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, tech_stack_id)
);