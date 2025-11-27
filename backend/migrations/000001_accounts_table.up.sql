CREATE TABLE accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    email TEXT NOT NULL UNIQUE CHECK (POSITION('@' IN email) > 1),
    hashed_password VARCHAR(72) NOT NULL,
    role TEXT CHECK (
        role IN (
            'user',
            'organizer',
            'employee',
            'admin'
        )
    ),
    disabled BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp
);