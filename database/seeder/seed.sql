-- Roles tables
INSERT INTO roles (name, description) VALUES 
('Admin', 'Administrator utama yang memiliki akses penuh ke seluruh sistem.'),
('Mahasiswa', 'Mahasiswa aktif yang dapat mengajukan klaim poin prestasi.'),
('Dosen Wali', 'Dosen pembimbing yang bertugas memverifikasi validitas prestasi mahasiswa.');

-- Permissions tables
INSERT INTO permissions (name, resource, action, description) VALUES 
('users:read',          'users',        'read',   'Melihat daftar pengguna/mahasiswa/dosen'),
('users:create',        'users',        'create', 'Membuat user baru'),
('users:update',        'users',        'update', 'Mengedit user (termasuk assign role/advisor)'),
('users:delete',        'users',        'delete', 'Menghapus user'),
('students:read',       'students',     'read',   'Melihat data detail mahasiswa'),
('lecturers:read',      'lecturers',    'read',   'Melihat data detail dosen'),
('achievements:read',   'achievements', 'read',   'Melihat daftar prestasi (Milik sendiri/Bimbingan)'),
('achievements:create', 'achievements', 'create', 'Membuat draft prestasi baru'),
('achievements:update', 'achievements', 'update', 'Mengedit prestasi (hanya status Draft)'),
('achievements:delete', 'achievements', 'delete', 'Menghapus prestasi (hanya status Draft)'),
('achievements:submit', 'achievements', 'submit', 'Mengirim prestasi untuk diverifikasi'),
('achievements:verify', 'achievements', 'verify', 'Menyetujui prestasi mahasiswa (Status: Verified)'),
('achievements:reject', 'achievements', 'reject', 'Menolak prestasi mahasiswa (Status: Rejected)'),
('reports:read',        'reports',      'read',   'Melihat dashboard statistik prestasi');

-- Insert User: George Admin
INSERT INTO users (username, email, password_hash, full_name, role_id) VALUES 
(
    'george_admin', 
    'george@admin.unair.ac.id', 
    '$2a$10$EixZaYVK1fsbw1ZfbX3OXePaWrn95nPn6as79w0hNtGOlqAttiikO', 
    'George Administrator',
    (SELECT id FROM roles WHERE name = 'Admin' LIMIT 1)
);