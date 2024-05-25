-- Insert roles
-- Create a unique constraint on the id column
INSERT IGNORE INTO roles (id, name, description, created_at, updated_at, deleted_at)
VALUES (1, 'Admin', 'Administrator role', NOW(), NOW(), NULL),
       (2, 'Manager', 'Manager role', NOW(), NOW(), NULL),
       (3, 'Employee', 'Employee role', NOW(), NOW(), NULL);