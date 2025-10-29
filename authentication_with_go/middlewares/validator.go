package middlewares

import (
	"AuthInGo/dto"
	"AuthInGo/utils"
	"context"
	"fmt"
	"net/http"
)

func UserLoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload *dto.LoginUserRequestDTO

		if err := utils.ReadJSONRequest(r, &payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid JSON request payload", err)
			return
		}

		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Validation failed for the request payload", err)
			return
		}

		fmt.Println("Payload received for login validation:", payload)

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}

func UserCreateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload *dto.CreateUserRequestDTO

		if err := utils.ReadJSONRequest(r, &payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid JSON request payload", err)
			return
		}

		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Validation failed for the request payload", err)
			return
		}

		fmt.Println("Payload received for user creation validation:", payload)

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}

func CreateRoleRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload *dto.CreateRoleRequestDTO
		if err := utils.ReadJSONRequest(r, &payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid JSON request payload", err)
			return
		}
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Validation failed for the request payload", err)
			return
		}
		fmt.Println("Payload received for role creation validation:", payload)

		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}

func UpdateRoleRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload *dto.UpdateRoleRequestDTO
		if err := utils.ReadJSONRequest(r, &payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid JSON request payload", err)
			return
		}
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Validation failed for the request payload", err)
			return
		}
		fmt.Println("Payload received for role update validation:", payload)

		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}

func AssignPermissionRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload *dto.AssignPermissionsRequestDTO
		if err := utils.ReadJSONRequest(r, &payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid JSON request payload", err)
			return
		}
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Validation failed for the request payload", err)
			return
		}
		fmt.Println("Payload received for permission assignment validation:", payload)

		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}

func RemovePermissionRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload *dto.RemovePermissionsRequestDTO
		if err := utils.ReadJSONRequest(r, &payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid JSON request payload", err)
			return
		}
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Validation failed for the request payload", err)
			return
		}
		fmt.Println("Payload received for permission removal validation:", payload)

		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}

func AssignRoleToUserRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload *dto.AssignRoleToUserRequestDTO
		if err := utils.ReadJSONRequest(r, &payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid JSON request payload", err)
			return
		}
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Validation failed for the request payload", err)
			return
		}
		fmt.Println("Payload received for role assignment validation:", payload)

		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}
