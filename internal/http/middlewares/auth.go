package middlewares

import "net/http"

// AuthMiddleware checks if user is authenticated
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Add your token validation logic here
		// You might want to parse JWT, check session, etc.

		next.ServeHTTP(w, r)
	})
}

// PermissionMiddleware checks if user has required role
func PermissionMiddleware(requiredRole string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get user role from context (usually set by AuthMiddleware)
			// This is just an example - adjust based on your auth implementation
			userRole := r.Context().Value("userRole")

			if userRole != requiredRole {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
