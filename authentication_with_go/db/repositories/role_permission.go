package db

import (
	"AuthInGo/models"
	"database/sql"
)

type RolePermissionRepository interface {
	GetRolePermissionByID(rolePermissionId int64) (*models.RolePermission, error)
	GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermission, error)
	AddPermissionToRole(roleId int64, permissionId int64) error
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	GetAllRolePermissions() ([]*models.RolePermission, error)
}

type RolePermissionRepositoryImpl struct {
	db *sql.DB
}

func NewRolePermissionRepository(_db *sql.DB) RolePermissionRepository {
	return &RolePermissionRepositoryImpl{
		db: _db,
	}
}

func (r *RolePermissionRepositoryImpl) GetRolePermissionByID(rolePermissionId int64) (*models.RolePermission, error) {
	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions WHERE id = ?"
	row := r.db.QueryRow(query, rolePermissionId)

	var rolePermission models.RolePermission
	if err := row.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId, &rolePermission.CreatedAt, &rolePermission.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &rolePermission, nil
}

func (r *RolePermissionRepositoryImpl) GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermission, error) {
	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions WHERE role_id = ?"
	rows, err := r.db.Query(query, roleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermission
	for rows.Next() {
		var rolePermission models.RolePermission
		if err := rows.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId, &rolePermission.CreatedAt, &rolePermission.UpdatedAt); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, &rolePermission)
	}
	return rolePermissions, nil
}

func (r *RolePermissionRepositoryImpl) AddPermissionToRole(roleId int64, permissionId int64) error {
	query := "INSERT INTO role_permissions (role_id, permission_id, created_at, updated_at) VALUES (?, ?, NOW(), NOW())"
	_, err := r.db.Exec(query, roleId, permissionId)
	return err
}

func (r *RolePermissionRepositoryImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {
	query := "DELETE FROM role_permissions WHERE role_id = ? AND permission_id = ?"
	_, err := r.db.Exec(query, roleId, permissionId)
	return err
}

func (r *RolePermissionRepositoryImpl) GetAllRolePermissions() ([]*models.RolePermission, error) {
	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermission
	for rows.Next() {
		var rolePermission models.RolePermission
		if err := rows.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId, &rolePermission.CreatedAt, &rolePermission.UpdatedAt); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, &rolePermission)
	}
	return rolePermissions, nil
}