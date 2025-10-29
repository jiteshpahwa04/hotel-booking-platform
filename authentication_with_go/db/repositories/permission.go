package db

import (
	"AuthInGo/models"
	"database/sql"
)

type PermissionRepository interface {
	GetPermissionByID(permissionId int64) (*models.Permission, error)
	GetPermissionByName(permissionName string) (*models.Permission, error)
	GetAllPermissions() ([]*models.Permission, error)
	CreatePermission(permission *models.Permission) (int64, error)
	UpdatePermission(permission *models.Permission) error
	DeletePermission(permissionId int64) error
}

type PermissionRepositoryImpl struct {
	db *sql.DB
}

func NewPermissionRepository(_db *sql.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db: _db,
	}
}

func (r *PermissionRepositoryImpl) GetPermissionByID(permissionId int64) (*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE id = ?"
	row := r.db.QueryRow(query, permissionId)
	var permission models.Permission
	if err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &permission, nil
}

func (r *PermissionRepositoryImpl) GetPermissionByName(permissionName string) (*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE name = ?"
	row := r.db.QueryRow(query, permissionName)
	var permission models.Permission
	if err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &permission, nil
}

func (r *PermissionRepositoryImpl) GetAllPermissions() ([]*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission
	for rows.Next() {
		var permission models.Permission
		if err := rows.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt); err != nil {
			return nil, err
		}
		permissions = append(permissions, &permission)
	}
	return permissions, nil
}

func (r *PermissionRepositoryImpl) CreatePermission(permission *models.Permission) (int64, error) {
	query := "INSERT INTO permissions (name, description, resource, action, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())"
	result, err := r.db.Exec(query, permission.Name, permission.Description, permission.Resource, permission.Action)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *PermissionRepositoryImpl) UpdatePermission(permission *models.Permission) error {
	query := "UPDATE permissions SET name = ?, description = ?, resource = ?, action = ?, updated_at = NOW() WHERE id = ?"
	_, err := r.db.Exec(query, permission.Name, permission.Description, permission.Resource, permission.Action, permission.Id)
	return err
}

func (r *PermissionRepositoryImpl) DeletePermission(permissionId int64) error {
	query := "DELETE FROM permissions WHERE id = ?"
	_, err := r.db.Exec(query, permissionId)
	return err
}