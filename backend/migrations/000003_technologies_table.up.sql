CREATE TABLE technologies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(150) NOT NULL UNIQUE
);

INSERT INTO
    technologies (name)
VALUES ('Go'),
    ('React'),
    ('Next.js'),
    ('C'),
    ('JavaScript'),
    ('TypeScript'),
    ('Node.js'),
    ('Python'),
    ('Java'),
    ('C#'),
    ('C++'),
    ('Rust'),
    ('PHP'),
    ('Ruby'),
    ('Swift'),
    ('Kotlin'),
    ('Dart'),
    ('Angular'),
    ('Vue')
ON CONFLICT (name) DO NOTHING;