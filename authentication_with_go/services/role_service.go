package services

import (
	repositories "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
)

type RoleService interface {
	GetRoleByID(roleId int64) (*models.Role, error)
	GetRoleByName(roleName string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(role *dto.CreateRoleRequestDTO) (int64, error)
	UpdateRole(role *dto.UpdateRoleRequestDTO, roleId int64) error
	DeleteRole(roleId int64) error
	GetRolePermissions(roleId int64) ([]*models.RolePermission, error)
	GetAllRolePermissions() ([]*models.RolePermission, error)
	AddPermissionToRole(roleId int64, permissionId int64) error
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	AssignRoleToUser(userId int64, roleId int64) error
}

type RoleServiceImpl struct {
	RoleRepository         repositories.RoleRepository
	RolePermissionRepository repositories.RolePermissionRepository
	UserRoleRepository    repositories.UserRoleRepository
}

func NewRoleService(roleRepository repositories.RoleRepository, rolePermissionRepository repositories.RolePermissionRepository, userRoleRepository repositories.UserRoleRepository) RoleService {
	return &RoleServiceImpl{
		RoleRepository:         roleRepository,
		RolePermissionRepository: rolePermissionRepository,
		UserRoleRepository:    userRoleRepository,
	}
}

func (s *RoleServiceImpl) GetRoleByID(roleId int64) (*models.Role, error) {
	return s.RoleRepository.GetRoleByID(roleId)
}

func (s *RoleServiceImpl) GetRoleByName(roleName string) (*models.Role, error) {
	return s.RoleRepository.GetRoleByName(roleName)
}

func (s *RoleServiceImpl) GetAllRoles() ([]*models.Role, error) {
	return s.RoleRepository.GetAllRoles()
}

func (s *RoleServiceImpl) CreateRole(role *dto.CreateRoleRequestDTO) (int64, error) {
	return s.RoleRepository.CreateRole(role.Name, role.Description)
}

func (s *RoleServiceImpl) UpdateRole(role *dto.UpdateRoleRequestDTO, roleId int64) error {
	return s.RoleRepository.UpdateRole(role.Name, role.Description, roleId)
}

func (s *RoleServiceImpl) DeleteRole(roleId int64) error {
	return s.RoleRepository.DeleteRole(roleId)
}

func (s *RoleServiceImpl) GetRolePermissions(roleId int64) ([]*models.RolePermission, error) {
	return s.RolePermissionRepository.GetRolePermissionByRoleId(roleId)
}

func (s *RoleServiceImpl) GetAllRolePermissions() ([]*models.RolePermission, error) {
	return s.RolePermissionRepository.GetAllRolePermissions()
}

func (s *RoleServiceImpl) AddPermissionToRole(roleId int64, permissionId int64) error {
	return s.RolePermissionRepository.AddPermissionToRole(roleId, permissionId)
}

func (s *RoleServiceImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {
	return s.RolePermissionRepository.RemovePermissionFromRole(roleId, permissionId)
}

func (s *RoleServiceImpl) AssignRoleToUser(userId int64, roleId int64) error {
	return s.UserRoleRepository.AssignRoleToUser(userId, roleId)
}