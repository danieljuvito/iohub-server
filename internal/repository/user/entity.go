package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entity struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    Email    string             `bson:"email"`
    Password string             `bson:"password"`
}
