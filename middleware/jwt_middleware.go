// middleware/jwt_middleware.go
package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/jwk"
)

// Define the Cognito details
const (
	Region     = "ap-south-1"
	UserPoolID = "ap-south-1_VOtG3qDHB"
	CognitoURL = "https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json"
	Issuer     = "https://cognito-idp.%s.amazonaws.com/%s"
)

// JWTAuth middleware for AWS Cognito
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify that the token's signing algorithm is what you expect
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Fetch the JWKS for this User Pool
			jwksURL := fmt.Sprintf(CognitoURL, Region, UserPoolID)
			keySet, err := jwk.Fetch(r.Context(), jwksURL)
			if err != nil {
				log.Println("Failed to fetch JWKs", err)
				return nil, err
			}

			// Extract the kid (key ID) from the JWT header and use it to find the correct public key
			kid := token.Header["kid"].(string)
			key, found := keySet.LookupKeyID(kid)
			if !found {
				return nil, fmt.Errorf("unable to find key")
			}

			// Return the public key for validation
			var rawKey interface{}
			if err := key.Raw(&rawKey); err != nil {
				return nil, err
			}
			return rawKey, nil
		})

		// Handle token parsing error or invalid token
		if err != nil || !token.Valid {
			log.Println("Invalid token:", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// If token is valid, continue with the request
		next.ServeHTTP(w, r)
	})
}
