package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	roleController *controllers.RoleController
}

func NewRoleRouter(_roleController *controllers.RoleController) Router {
	return &RoleRouter{
		roleController: _roleController,
	}
}

func (rr *RoleRouter) Register(r chi.Router) {
	// Role CRUD operations
	r.Get("/roles/{roleId}", rr.roleController.GetRoleByID)
	r.Get("/roles", rr.roleController.GetAllRoles)
	r.With(middlewares.CreateRoleRequestValidator).Post("/roles", rr.roleController.CreateRole)
	r.With(middlewares.UpdateRoleRequestValidator).Put("/roles/{roleId}", rr.roleController.UpdateRole)
	r.Delete("/roles/{roleId}", rr.roleController.DeleteRole)

	// Role Permission operations
	r.Get("/roles/{roleId}/permissions", rr.roleController.GetRolePermissions)
	r.With(middlewares.AssignPermissionRequestValidator).Post("/roles/{roleId}/permissions", rr.roleController.AddPermissionToRole)
	r.With(middlewares.RemovePermissionRequestValidator).Delete("/roles/{roleId}/permissions", rr.roleController.RemovePermissionFromRole)
	r.Get("/role-permissions", rr.roleController.GetAllRolePermissions)

	// User Role operations
	r.With(middlewares.JWTAuthMiddleware, middlewares.AssignRoleToUserRequestValidator, middlewares.RequireAllRoles(1)).Post("/assign-role", rr.roleController.AssignRoleToUser)
}