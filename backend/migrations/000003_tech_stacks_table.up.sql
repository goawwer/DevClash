CREATE TABLE tech_stacks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(150) NOT NULL UNIQUE
);

INSERT INTO
    tech_stacks (name)
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