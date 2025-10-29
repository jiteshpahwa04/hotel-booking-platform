-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    resource VARCHAR(255) NOT NULL,
    action VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- seeder data
-- INSERT INTO permissions (name, description, resource, action) VALUES
-- ('user:read', 'Permission to view user list', 'users', 'read'),
-- ('user:write', 'Permission to edit user details', 'users', 'write'),
-- ('user:delete', 'Permission to delete users', 'users', 'delete'),
-- ('role:read', 'Permission to view role list', 'roles', 'read'),
-- ('role:write', 'Permission to edit role details', 'roles', 'write'),
-- ('role:delete', 'Permission to delete roles', 'roles', 'delete'),
-- ('role:manage', 'Permission to manage role assignments', 'roles', 'manage'),
-- ('permission:read', 'Permission to view permissions', 'permissions', 'read'),
-- ('permission:write', 'Permission to edit permissions', 'permissions', 'write'),
-- ('permission:delete', 'Permission to delete permissions', 'permissions', 'delete'),
-- ('permission:manage', 'Permission to manage permissions', 'permissions', 'manage');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
