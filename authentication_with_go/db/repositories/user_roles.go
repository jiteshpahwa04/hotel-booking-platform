package db

import (
	"AuthInGo/models"
	"database/sql"
	"strings"
)

type UserRoleRepository interface {
	GetUserRoles(userId int64) ([]*models.Role, error)
	AssignRoleToUser(userId int64, roleId int64) error
	AddUserRole(userId int64, roleId int64) error
	RemoveUserRole(userId int64, roleId int64) error
	GetUserPermissions(userId int64) ([]*models.Permission, error)
	HasPermission(userId int64, permissionId int64) (bool, error)
	HasRole(userId int64, roleId int64) (bool, error)
	HasAllRoles(userId int64, roleIds []int64) (bool, error)
	HasAnyRole(userId int64, roleIds []int64) (bool, error)
}

type UserRoleRepositoryImpl struct {
	db *sql.DB
}

func NewUserRoleRepository(_db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImpl{
		db: _db,
	}
}

func (r *UserRoleRepositoryImpl) GetUserRoles(userId int64) ([]*models.Role, error) {
	query := `
	SELECT r.id, r.name, r.description, r.created_at, r.updated_at
	FROM roles r
	JOIN user_roles ur ON r.id = ur.role_id
	WHERE ur.user_id = ?`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func (r *UserRoleRepositoryImpl) AddUserRole(userId int64, roleId int64) error {
	query := "INSERT INTO user_roles (user_id, role_id, created_at, updated_at) VALUES (?, ?, NOW(), NOW())"
	_, err := r.db.Exec(query, userId, roleId)
	return err
}

func (r *UserRoleRepositoryImpl) RemoveUserRole(userId int64, roleId int64) error {
	query := "DELETE FROM user_roles WHERE user_id = ? AND role_id = ?"
	_, err := r.db.Exec(query, userId, roleId)
	return err
}

func (r *UserRoleRepositoryImpl) GetUserPermissions(userId int64) ([]*models.Permission, error) {
	query := `
	SELECT p.id, p.name, p.description, p.resource, p.action, p.created_at, p.updated_at
	FROM permissions p
	JOIN role_permissions rp ON p.id = rp.permission_id
	JOIN user_roles ur ON rp.role_id = ur.role_id
	WHERE ur.user_id = ?`
	rows, err := r.db.Query(query, userId)
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

func (r *UserRoleRepositoryImpl) HasPermission(userId int64, permissionId int64) (bool, error) {
	query := `
	SELECT COUNT(*)
	FROM user_roles ur
	JOIN role_permissions rp ON ur.role_id = rp.role_id
	WHERE ur.user_id = ? AND rp.permission_id = ?`
	var count int
	if err := r.db.QueryRow(query, userId, permissionId).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRoleRepositoryImpl) HasRole(userId int64, roleId int64) (bool, error) {
	query := "SELECT COUNT(*) FROM user_roles WHERE user_id = ? AND role_id = ?"
	var count int
	if err := r.db.QueryRow(query, userId, roleId).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRoleRepositoryImpl) HasAllRoles(userId int64, roleIds []int64) (bool, error) {
	if len(roleIds) == 0 {
		return true, nil
	}

	query := `
		SELECT COUNT(*) = ?
		FROM user_roles ur
		WHERE ur.user_id = ? AND ur.role_id IN (` + strings.TrimRight(strings.Repeat("?,", len(roleIds)), ",") + `)`
	
	var count int
	args := make([]interface{}, 0, len(roleIds)+2)
	args = append(args, len(roleIds), userId)
	for _, roleId := range roleIds {
		args = append(args, roleId)
	}
	if err := r.db.QueryRow(query, args...).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRoleRepositoryImpl) HasAnyRole(userId int64, roleIds []int64) (bool, error) {
	if len(roleIds) == 0 {
		return false, nil
	}

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM user_roles ur
			WHERE ur.user_id = ? AND ur.role_id IN (` + strings.TrimRight(strings.Repeat("?,", len(roleIds)), ",") + `)
		)`

	var hasRole bool
	args := make([]interface{}, 0, len(roleIds)+1)
	args = append(args, userId) // because we cannot use userId in the IN clause with a variable number of arguments
	for _, roleId := range roleIds {
		args = append(args, roleId)
	}
	if err := r.db.QueryRow(query, args...).Scan(&hasRole); err != nil {
		return false, err
	}

	return hasRole, nil
}

func (r *UserRoleRepositoryImpl) AssignRoleToUser(userId int64, roleId int64) error {
	// First, check if the user already has the role
	hasRole, err := r.HasRole(userId, roleId)
	if err != nil {
		return err
	}
	if hasRole {
		// User already has the role, no need to assign
		return nil
	}
	// Assign the role to the user
	return r.AddUserRole(userId, roleId)
}