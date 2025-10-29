package db

import (
	"AuthInGo/models"
	"database/sql"
)

type RoleRepository interface {
	GetRoleByID(roleId int64) (*models.Role, error)
	GetRoleByName(roleName string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(name string, description string) (int64, error)
	UpdateRole(name string, description string, roleId int64) error
	DeleteRole(roleId int64) error
}

type RoleRepositoryImpl struct {
	db *sql.DB
}

func NewRoleRepository(_db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: _db,
	}
}

func (r *RoleRepositoryImpl) GetRoleByID(roleId int64) (*models.Role, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE id = ?"
	row := r.db.QueryRow(query, roleId)

	var role models.Role
	if err := row.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepositoryImpl) GetRoleByName(roleName string) (*models.Role, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE name = ?"
	row := r.db.QueryRow(query, roleName)

	var role models.Role
	if err := row.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepositoryImpl) GetAllRoles() ([]*models.Role, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM roles"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // defers the closing of rows after function ends

	var roles []*models.Role
	for rows.Next() {
		var role models.Role
		if err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}
	return roles, nil
}

func (r *RoleRepositoryImpl) CreateRole(name string, description string) (int64, error) {
	query := "INSERT INTO roles (name, description, created_at, updated_at) VALUES (?, ?, NOW(), NOW())"
	result, err := r.db.Exec(query, name, description)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *RoleRepositoryImpl) UpdateRole(name string, description string, roleId int64) error {
	query := "UPDATE roles SET name = ?, description = ?, updated_at = NOW() WHERE id = ?"
	_, err := r.db.Exec(query, name, description, roleId)
	return err
}

func (r *RoleRepositoryImpl) DeleteRole(roleId int64) error {
	query := "DELETE FROM roles WHERE id = ?"
	_, err := r.db.Exec(query, roleId)
	return err
}