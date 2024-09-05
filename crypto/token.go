package crypto

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	// "time"
)

const (
	// Secret key used for signing the JWT token
	secretKey = "your-secret-key"
)

func GenerateToken(claims jwt.StandardClaims) (string, error) {
	// Create the JWT token with a custom set of claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	// Parse the JWT token and verify the signature
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		// Check if the token parsing error is due to the wrong signing method
		if _, ok := err.(*jwt.ValidationError); ok {
			// Try parsing the token with a different signing method
			token, err = jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// Extract the custom claims from the token
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
