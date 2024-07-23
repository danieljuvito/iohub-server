package story

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Entity struct {
    ID        primitive.ObjectID
    UserID    primitive.ObjectID
    Content   string
    ExpiresAt time.Time
}
