package middlewares

import (
	"context"
	"net/http"

	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/token"
	"github.com/midedickson/instashop/utils"
)

// AuthMiddleware checks if user is authenticated
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := token.ExtractFromHeader(r)
		if accessToken == "" {
			utils.Dispatch403Error(w, "Invalid or Expired token", nil)
			return
		}

		tokenValid, claim, errTokenVerify := token.Verify(&token.TokenVerifyOptions{SignedToken: accessToken})
		if errTokenVerify != nil || !tokenValid {
			utils.Dispatch403Error(w, "Invalid or Expired token", errTokenVerify)
			return
		}

		ctx := context.WithValue(r.Context(), constants.AuthClaimCtxKey{}, claim.Payload)
		ctx = context.WithValue(ctx, constants.UserRoleCtxKey{}, claim.Payload["role"])
		ctx = context.WithValue(ctx, constants.UserEmailCtxKey{}, claim.Payload["email"])
		ctx = context.WithValue(ctx, constants.UserIdCtxKey{}, claim.Payload["id"])

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// PermissionMiddleware checks if user has required role
func PermissionMiddleware(requiredRole string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := r.Context().Value(constants.UserRoleCtxKey{}).(string)

			if userRole != requiredRole {
				utils.Dispatch403Error(w, "Invalid or Expired token", utils.ErrForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
