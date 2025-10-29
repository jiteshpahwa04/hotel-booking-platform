-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS role_permissions (
    id SERIAL PRIMARY KEY,
    role_id BIGINT UNSIGNED NOT NULL,
    permission_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);

-- seeder data
-- INSERT INTO role_permissions (role_id, permission_id)
-- SELECT 1, id FROM permissions; -- Assuming admin has role id 1, admin gets all permissions
-- INSERT INTO role_permissions (role_id, permission_id)
-- SELECT 2, id FROM permissions WHERE name IN ('user:read', 'user:write'); -- user role with limited permissions
-- INSERT INTO role_permissions (role_id, permission_id)
-- SELECT 3, id FROM permissions WHERE name IN ('user:read', 'role:read'); -- moderator role with minimal permissions
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role_permissions;
-- +goose StatementEnd
