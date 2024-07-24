package middleware

import (
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "github.com/danieljuvito/iohub-server/internal/repository/session"
    "github.com/danieljuvito/iohub-server/internal/util/jwtutil"
    "github.com/golang-jwt/jwt/v5"
    "github.com/labstack/echo/v4"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "net/http"
    "strings"
    "time"
)

var db *mongo.Database

func InitAuth(_db *mongo.Database) {
    db = _db
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        ctx := c.Request().Context()

        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return c.JSON(http.StatusUnauthorized, "unauthorized")
        }

        tokenString = strings.Split(tokenString, " ")[1]
        token, err := jwtutil.VerifyToken(tokenString)
        if err != nil {
            return c.JSON(http.StatusUnauthorized, "invalid token")
        }

        claims := token.Claims.(jwt.MapClaims)

        now := time.Now()
        exp := claims["exp"].(float64)
        if float64(now.Unix())-exp > 0 {
            return c.JSON(http.StatusUnauthorized, "session is expired")
        }

        sessionID := claims["sub"].(string)
        objectID, err := primitive.ObjectIDFromHex(sessionID)
        if err != nil {
            return c.JSON(http.StatusUnauthorized, "invalid session")
        }

        entity := new(session.Entity)
        err = db.Collection("sessions").FindOne(ctx, bson.M{"_id": objectID}).Decode(entity)
        if err != nil || entity == nil {
            return c.JSON(http.StatusUnauthorized, "session not found")
        }

        if now.Sub(entity.ExpiresAt) > 0 {
            return c.JSON(http.StatusUnauthorized, "session is expired")
        }

        c.Set("session", model.Session{
            ID:        entity.ID.Hex(),
            UserID:    entity.UserID.Hex(),
            ExpiresAt: entity.ExpiresAt,
        })

        return next(c)
    }
}
