package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// UserContextKey adalah tipe khusus untuk key context
type UserContextKey string

const (
	UserIDKey   UserContextKey = "user_id"
	UserRoleKey UserContextKey = "user_role"
)

// RequireAuth memproteksi endpoint agar hanya bisa diakses dengan token JWT yang valid.
func RequireAuth(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				respondError(w, http.StatusUnauthorized, "Authorization header is required")
				return
			}

			// Format header harus "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				respondError(w, http.StatusUnauthorized, "Invalid authorization format")
				return
			}

			tokenString := parts[1]

			// Parse token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Pastikan algoritma yang digunakan sesuai
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secretKey), nil
			})

			if err != nil || !token.Valid {
				respondError(w, http.StatusUnauthorized, "Invalid or expired token")
				return
			}

			// Ekstrak klaim (user_id dan role)
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				var userID uint
				var userRole string

				if userIDFloat, ok := claims["user_id"].(float64); ok {
					userID = uint(userIDFloat)
				} else {
					respondError(w, http.StatusUnauthorized, "Invalid token claims: missing user_id")
					return
				}

				if roleStr, ok := claims["role"].(string); ok {
					userRole = roleStr
				}

				// Masukkan user ID dan Role ke dalam context request
				ctx := context.WithValue(r.Context(), UserIDKey, userID)
				ctx = context.WithValue(ctx, UserRoleKey, userRole)

				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			respondError(w, http.StatusUnauthorized, "Invalid token claims")
		})
	}
}

// RequireRole memeriksa apakah user role yang tersimpan di context cocok dengan daftar allowed roles.
// HARUS dipanggil setelah RequireAuth.
func RequireRole(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole, ok := r.Context().Value(UserRoleKey).(string)
			if !ok || userRole == "" {
				respondError(w, http.StatusForbidden, "Access denied: missing role")
				return
			}

			roleAllowed := false
			for _, role := range allowedRoles {
				if userRole == role {
					roleAllowed = true
					break
				}
			}

			if !roleAllowed {
				respondError(w, http.StatusForbidden, "Access denied: insufficient permissions")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func respondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
