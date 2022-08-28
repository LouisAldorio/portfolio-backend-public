package utils

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type UserClaim struct {
	Username string
	jwt.StandardClaims
}

func Auth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authToken := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if authToken == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			jwtToken, err := ValidateToken(authToken)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			//validate claim
			claims, ok := jwtToken.Claims.(*UserClaim)
			if !ok && !jwtToken.Valid {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			//return user data to req
			ctx := context.WithValue(r.Context(), userCtxKey, claims)
			reqWithCtx := r.WithContext(ctx)
			next.ServeHTTP(w, reqWithCtx)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func AuthContext(ctx context.Context) *UserClaim {
	raw, _ := ctx.Value(userCtxKey).(*UserClaim)
	return raw
}
