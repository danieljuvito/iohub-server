package jwtutil

import (
    "github.com/golang-jwt/jwt/v5"
    "os"
    "time"

    //    "time"
)

var secretKey = os.Getenv("JWT_KEY")

type Claims struct {
    Name string `json:"name"`
    jwt.RegisteredClaims
}

func GenerateToken(sessionID string, exp time.Time) (string, error) {
    claims := jwt.MapClaims{
        "sub": sessionID,
        "exp": exp.Unix(), // Token expires in 24 hours
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
    })
}
