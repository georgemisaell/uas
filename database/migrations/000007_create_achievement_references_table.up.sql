-- 1. Membuat Tipe Data ENUM (Jalankan ini dulu)
-- DO block ini digunakan untuk mengecek apakah type sudah ada agar tidak error jika dijalankan ulang
DO $$ BEGIN
    CREATE TYPE achievement_status_enum AS ENUM ('draft', 'submitted', 'verified', 'rejected');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- 2. Membuat Tabel achievement_references
CREATE TABLE IF NOT EXISTS achievement_references (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    student_id UUID NOT NULL,
    mongo_achievement_id VARCHAR(24) NOT NULL,
    status achievement_status_enum NOT NULL DEFAULT 'draft',
    submitted_at TIMESTAMP,
    verified_at TIMESTAMP,
    verified_by UUID,
    rejection_note TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_achievement_student
        FOREIGN KEY (student_id)
        REFERENCES students(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_achievement_verifier
        FOREIGN KEY (verified_by)
        REFERENCES users(id)
        ON DELETE SET NULL
);