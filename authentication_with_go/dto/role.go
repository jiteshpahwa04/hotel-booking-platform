package dto

type CreateRoleRequestDTO struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

type UpdateRoleRequestDTO struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

type AssignPermissionsRequestDTO struct {
	PermissionId int64 `json:"permissionId"`
}

type RemovePermissionsRequestDTO struct {
	PermissionId int64 `json:"permissionId"`
}

type AssignRoleToUserRequestDTO struct {
	UserId int64 `json:"userId"`
	RoleId int64 `json:"roleId"`
}