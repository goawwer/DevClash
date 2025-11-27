CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    account_id UUID REFERENCES accounts (id) ON DELETE CASCADE,
    username VARCHAR(64) NOT NULL UNIQUE CHECK (
        LENGTH(username) >= 5
        AND LENGTH(username) <= 64
    ),
    profile_picture_url TEXT,
    bio TEXT,
    profile_status VARCHAR(150),
    participations_count INT NOT NULL DEFAULT 0,
    wins_count INT NOT NULL DEFAULT 0
);