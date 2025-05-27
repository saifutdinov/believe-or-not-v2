package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// JwtSecretKey should be set from environment for production
	JwtSecretKey    = "CHANGE_ME_SECRET"
	JwtExpirePeriod = time.Hour * 3
)

type contextKey string

const userContextKey contextKey = "userID"

// GenerateJWT создает JWT для заданного userID
func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(JwtExpirePeriod).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JwtSecretKey))
}

// JWTMiddleware проверяет заголовок Authorization: Bearer <token>
// Валидирует токен и кладет userID в контекст
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]
		tok, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(JwtSecretKey), nil
		})
		if err != nil || !tok.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := tok.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		sub, ok := claims["sub"].(float64)
		if !ok {
			http.Error(w, "Invalid subject claim", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, sub)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserIDFromContext извлекает userID из контекста
func UserIDFromContext(ctx context.Context) (float64, bool) {
	sub, ok := ctx.Value(userContextKey).(float64)
	return sub, ok
}
