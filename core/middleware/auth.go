package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "secret"

type contextKey string

const userKey contextKey = "user"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow OPTIONS requests
		if r.Method == "OPTIONS" || r.Method == "GET" {
			next.ServeHTTP(w, r)
			return
		}

		// Read the body once, then replace it so other handlers can read it too
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		r.Body.Close()                                    // Close the original body
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Replace with new readable body

		var reqBody struct {
			Query         string `json:"query"`
			OperationName string `json:"operationName"` // Clients should send this along with the query
		}
		if err := json.Unmarshal(bodyBytes, &reqBody); err == nil {

			fmt.Println(reqBody.Query)

			// Check operation name or presence of auth-related operations
			if reqBody.OperationName == "Login" || reqBody.OperationName == "Register" || reqBody.OperationName == "IntrospectionQuery" {
				next.ServeHTTP(w, r)
				return
			}
		}
		// Check if the request is a GraphQL query
		if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "Invalid Content-Type. Only application/json is allowed", http.StatusUnsupportedMediaType)
			return
		}

		// Process the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SecretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Attach user information to the context
		ctx := context.WithValue(r.Context(), userKey, token.Claims.(jwt.MapClaims))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
