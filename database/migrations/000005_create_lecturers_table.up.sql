CREATE TABLE lecturers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    lecturer_id VARCHAR(20),
    department VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW()
);