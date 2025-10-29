package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoleController struct {
	RoleService services.RoleService
}

func NewRoleController(roleService services.RoleService) *RoleController {
	return &RoleController{
		RoleService: roleService,
	}
}

func (rc *RoleController) GetRoleByID(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "roleId")
	if roleId == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role id"))
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", fmt.Errorf("invalid role id"))
		return
	}

	role, err := rc.RoleService.GetRoleByID(id)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to get role", err)
		return
	}

	if role == nil {
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "Role not found", fmt.Errorf("role not found"))
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role fetched successfully", role)
}

func (rc *RoleController) GetRoleByName(w http.ResponseWriter, r *http.Request) {
	roleName := r.URL.Query().Get("name")
	if roleName == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role name is required", fmt.Errorf("missing role name"))
		return
	}

	role, err := rc.RoleService.GetRoleByName(roleName)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to get role", err)
		return
	}

	if role == nil {
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "Role not found", fmt.Errorf("role not found"))
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role fetched successfully", role)
}

func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := rc.RoleService.GetAllRoles()
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to get roles", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Roles fetched successfully", roles)
}

func (rc *RoleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value("payload").(*dto.CreateRoleRequestDTO)

	roleId, err := rc.RoleService.CreateRole(payload); 
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to create role", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusCreated, "Role created successfully", roleId)
}

func (rc *RoleController) UpdateRole(w http.ResponseWriter, r *http.Request) {
	roleIdStr := chi.URLParam(r, "roleId")
	if roleIdStr == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role id"))
		return
	}

	roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", fmt.Errorf("invalid role id"))
		return
	}

	payload := r.Context().Value("payload").(*dto.UpdateRoleRequestDTO)
	if err := rc.RoleService.UpdateRole(payload, roleId); err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to update role", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role updated successfully", roleId)
}

func (rc *RoleController) DeleteRole(w http.ResponseWriter, r *http.Request) {
	roleIdStr := chi.URLParam(r, "roleId")
	if roleIdStr == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role id"))
		return
	}

	roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", fmt.Errorf("invalid role id"))
		return
	}

	if err := rc.RoleService.DeleteRole(roleId); err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to delete role", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role deleted successfully", roleId)
}

func (rc *RoleController) GetRolePermissions(w http.ResponseWriter, r *http.Request) {
	roleIdStr := chi.URLParam(r, "roleId")
	if roleIdStr == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role id"))
		return
	}

	roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", fmt.Errorf("invalid role id"))
		return
	}

	permissions, err := rc.RoleService.GetRolePermissions(roleId)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to get role permissions", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role permissions fetched successfully", permissions)
}

func (rc *RoleController) GetAllRolePermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := rc.RoleService.GetAllRolePermissions()
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to get role permissions", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role permissions fetched successfully", permissions)
}

func (rc *RoleController) AddPermissionToRole(w http.ResponseWriter, r *http.Request) {
	roleIdStr := chi.URLParam(r, "roleId")
	if roleIdStr == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role id"))
		return
	}

	roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", fmt.Errorf("invalid role id"))
		return
	}

	payload := r.Context().Value("payload").(*dto.AssignPermissionsRequestDTO)
	if err := rc.RoleService.AddPermissionToRole(roleId, payload.PermissionId); err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to add permission to role", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Permission added to role successfully", roleId)
}

func (rc *RoleController) RemovePermissionFromRole(w http.ResponseWriter, r *http.Request) {
	roleIdStr := chi.URLParam(r, "roleId")
	if roleIdStr == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role id"))
		return
	}

	roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", fmt.Errorf("invalid role id"))
		return
	}

	payload := r.Context().Value("payload").(*dto.RemovePermissionsRequestDTO)
	if err := rc.RoleService.RemovePermissionFromRole(roleId, payload.PermissionId); err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to remove permission from role", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Permission removed from role successfully", roleId)
}

func (rc *RoleController) AssignRoleToUser(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value("payload").(*dto.AssignRoleToUserRequestDTO)
	if err := rc.RoleService.AssignRoleToUser(payload.UserId, payload.RoleId); err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to assign role to user", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role assigned to user successfully", payload)
}