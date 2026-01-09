package middleware

import (
	"context"
	"net/http"
	"strings"
	"utils/logger"
)

type contextKey string

const (
	// BearerTokenKey is the context key for the bearer token
	BearerTokenKey contextKey = "bearer_token"
)

// TokenValidator is a function type for validating bearer tokens
type TokenValidator func(token string) (bool, error)

// BearerAuth returns a middleware that validates Bearer tokens
func BearerAuth(validate TokenValidator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				unauthorized(w, "missing authorization header")
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				unauthorized(w, "invalid authorization header format")
				return
			}

			token := parts[1]
			if token == "" {
				unauthorized(w, "empty bearer token")
				return
			}

			valid, err := validate(token)
			if err != nil {
				logger.Error("token validation error", "error", err)
				unauthorized(w, "token validation failed")
				return
			}

			if !valid {
				unauthorized(w, "invalid token")
				return
			}

			// Store token in context for downstream handlers
			ctx := context.WithValue(r.Context(), BearerTokenKey, token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func unauthorized(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`{"error":"` + message + `"}`))
}

// GetBearerToken retrieves the bearer token from context
func GetBearerToken(ctx context.Context) string {
	token, _ := ctx.Value(BearerTokenKey).(string)
	return token
}
