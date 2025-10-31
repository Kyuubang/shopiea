-- Seed data for Shopiea application
-- This file contains initial data for the database

-- Insert roles (admin and student)
INSERT INTO roles (id, name) VALUES 
    (1, 'admin'),
    (2, 'student')
ON CONFLICT (name) DO NOTHING;

-- Insert default class
INSERT INTO classes (id, name) VALUES 
    (1, 'Shopiea')
ON CONFLICT (name) DO NOTHING;

-- Insert default admin user
-- Password: admin123 (hashed with bcrypt)
INSERT INTO users (id, username, password, name, role_id, class_id) VALUES 
    (1, 'admin', '$2a$10$65liQ6mYNskzwSG9pj4qA.gXhpDwIeXXv5o0gb/dpr.uuCZymWXSm', 'Administrator', 1, 1)
ON CONFLICT (username) DO NOTHING;

-- Insert default course
INSERT INTO courses (id, name) VALUES 
    (1, 'golang')
ON CONFLICT (name) DO NOTHING;

-- Insert default lab
INSERT INTO labs (id, name, course_id) VALUES 
    (1, 'golang-001', 1)
ON CONFLICT (name) DO NOTHING;

-- Reset sequences to avoid conflicts
SELECT setval('roles_id_seq', (SELECT MAX(id) FROM roles));
SELECT setval('classes_id_seq', (SELECT MAX(id) FROM classes));
SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));
SELECT setval('courses_id_seq', (SELECT MAX(id) FROM courses));
SELECT setval('labs_id_seq', (SELECT MAX(id) FROM labs));
