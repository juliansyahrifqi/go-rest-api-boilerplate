package middlewares

import (
	"fmt"
	"go-rest-api-boilerplate/utils"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
)

// Middleware for Set A ContentType json
func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Auth Middleware
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromHeader(r)

		if tokenString == "" {
			utils.ResponseWithError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := validateToken(tokenString)

		if err != nil {
			utils.ResponseWithError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*jwt.MapClaims)

		if !ok {
			utils.ResponseWithError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claimsMap := make(map[string]interface{})

		for key, value := range *claims {
			claimsMap[key] = value
		}

		context.Set(r, "user", claimsMap)
		context.Set(r, "token", tokenString)

		next.ServeHTTP(w, r)
	})
}

func getTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return ""
	}

	tokenParts := strings.Split(authHeader, " ")

	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return ""
	}

	return tokenParts[1]
}

func validateToken(tokenString string) (*jwt.Token, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
