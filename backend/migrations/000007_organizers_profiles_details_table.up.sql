CREATE TABLE organizers_details (
    organizer_id UUID PRIMARY KEY REFERENCES organizers (id) ON DELETE CASCADE,
    company_description TEXT,
    logo_url TEXT,
    brand_color VARCHAR(30)
)