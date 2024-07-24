package story

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Entity struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    UserID    primitive.ObjectID `bson:"user_id"`
    Content   string             `bson:"content"`
    ExpiresAt time.Time          `bson:"expires_at"`
}
