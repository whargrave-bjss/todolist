package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"context"
	"errors"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing authorization header", http.StatusUnauthorized)
            return
        }

 
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        userID, err := validateTokenAndGetUserID(tokenString)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Set the user ID in the request context
        ctx := context.WithValue(r.Context(), "UserID", userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    }
}

func validateTokenAndGetUserID(tokenString string) (int, error) {
    secretKey := []byte("your-secret-key-here")

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return secretKey, nil
    })

    if err != nil {
        return 0, fmt.Errorf("error parsing token: %w", err)
    }

    if !token.Valid {
        return 0, errors.New("invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return 0, errors.New("invalid token claims")
    }

    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        return 0, errors.New("user ID not found in token")
    }

    userID := int(userIDFloat)

    expirationTime, ok := claims["exp"].(float64)
    if !ok {
        return 0, errors.New("expiration time not found in token")
    }

    if float64(time.Now().Unix()) > expirationTime {
        return 0, errors.New("token has expired")
    }

    return userID, nil
}

func CreateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
        "iat":     time.Now().Unix(),
    })

    return token.SignedString([]byte("your-secret-key-here"))
}