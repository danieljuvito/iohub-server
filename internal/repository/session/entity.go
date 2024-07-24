package session

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Entity struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    UserID    primitive.ObjectID `bson:"user_id"`
    ExpiresAt time.Time          `bson:"expires_at"`
}
