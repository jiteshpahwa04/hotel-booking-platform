package middlewares

import (
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
	repo "AuthInGo/db/repositories"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer") {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Token missing in authorization header", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetString("JWT_SECRET", "auth-in-go-secret")), nil
		})

		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		fmt.Println("Claims: ", claims)

		userId, ok := claims["id"].(float64)
		email, okEmail := claims["email"].(string)
		if !ok || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		fmt.Println("Authenticated user ID from token:", userId, "Email:", email)

		ctx := context.WithValue(r.Context(), "userId", strconv.FormatFloat(userId, 'f', 0, 64))
		ctx = context.WithValue(ctx, "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireAllRoles(roles ...int64) func(http.Handler) http.Handler {
	// function that can create a middleware to check the above set of roles
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userIdStr := r.Context().Value("userId").(string)
			userId, err := strconv.ParseInt(userIdStr, 10, 64)
			if err != nil {
				http.Error(w, "Invalid user ID in context", http.StatusUnauthorized)
				return
			}

			dbConn, err := dbConfig.SetupDb()
			if err != nil {
				http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
				return
			}
			defer dbConn.Close()

			urr := repo.NewUserRoleRepository(dbConn)

			hasAllRoles, err := urr.HasAllRoles(userId, roles)
			if err != nil {
				http.Error(w, "Error checking user roles", http.StatusInternalServerError)
				return
			}

			if !hasAllRoles {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func RequireAnyRole(roles ...int64) func(http.Handler) http.Handler {
	// function that can create a middleware to check the above set of roles
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userIdStr := r.Context().Value("userId").(string)
			userId, err := strconv.ParseInt(userIdStr, 10, 64)
			if err != nil {
				http.Error(w, "Invalid user ID in context", http.StatusUnauthorized)
				return
			}

			dbConn, err := dbConfig.SetupDb()
			if err != nil {
				http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
				return
			}
			defer dbConn.Close()

			urr := repo.NewUserRoleRepository(dbConn)

			hasAnyRole, err := urr.HasAnyRole(userId, roles)
			if err != nil {
				fmt.Println("Error checking user roles", err)
				http.Error(w, "Error checking user roles", http.StatusInternalServerError)
				return
			}

			if !hasAnyRole {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
