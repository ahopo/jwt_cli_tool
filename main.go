package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Generate a JWT token
func generateToken(key []byte, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

// Verify a JWT token
func verifyToken(tokenString string, key []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func main() {
	// Command-line arguments
	action := flag.String("action", "", "Action to perform: generate or verify")
	tokenString := flag.String("token", "", "Token to verify (required if action is 'verify')")
	secret := flag.String("secret", "", "Secret key for signing and verifying tokens (required)")
	subject := flag.String("sub", "", "Subject claim for token (required if action is 'generate')")
	name := flag.String("name", "", "Name claim for token (required if action is 'generate')")

	flag.Parse()

	if *action == "" || *secret == "" {
		fmt.Println("Error: action and secret are required")
		flag.Usage()
		return
	}

	secretKey := []byte(*secret)

	switch *action {
	case "generate":
		if *subject == "" || *name == "" {
			fmt.Println("Error: subject and name are required for generating a token")
			flag.Usage()
			return
		}

		// Example claims
		claims := jwt.MapClaims{
			"sub":  *subject,
			"name": *name,
			"iat":  time.Now().Unix(),
		}

		// Generate token
		token, err := generateToken(secretKey, claims)
		if err != nil {
			fmt.Println("Error generating token:", err)
			return
		}
		fmt.Println("Generated token:", token)

	case "verify":
		if *tokenString == "" {
			fmt.Println("Error: token is required for verification")
			flag.Usage()
			return
		}

		// Verify token
		token, err := verifyToken(*tokenString, secretKey)
		if err != nil {
			fmt.Println("Error verifying token:", err)
			return
		}
		fmt.Println("Token verified:", token)
	default:
		fmt.Println("Error: unknown action")
		flag.Usage()
	}
}
