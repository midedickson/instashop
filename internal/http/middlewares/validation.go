package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/midedickson/instashop/utils"
)

type Validatable interface {
	Validate() bool
}

// Middleware that validates the payload
func ValidatePayloadMiddleware(payload Validatable, ctxKey any) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Read and decode the body
			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
				utils.Dispatch400Error(w, "Invalid JSON payload", err)
				return
			}

			// Important: Restore the body for downstream handlers
			bodyBytes, _ := json.Marshal(payload)
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			// Validate the payload
			if !payload.Validate() {
				utils.Dispatch400Error(w, "Invalid payload", nil)
				return
			}

			ctx := context.WithValue(r.Context(), ctxKey, payload)

			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
