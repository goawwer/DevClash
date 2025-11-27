CREATE TABLE tech_stacks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(150) NOT NULL UNIQUE
);