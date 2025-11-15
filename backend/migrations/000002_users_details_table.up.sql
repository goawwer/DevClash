CREATE TABLE users_details (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    image_url TEXT,
    bio TEXT,
    profile_status VARCHAR(150)
);
