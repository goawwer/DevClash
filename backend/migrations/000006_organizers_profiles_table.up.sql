CREATE TABLE organizers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    account_id UUID REFERENCES accounts (id) ON DELETE CASCADE,
    name VARCHAR(150) NOT NULL UNIQUE,
    is_verified BOOLEAN DEFAULT FALSE
)