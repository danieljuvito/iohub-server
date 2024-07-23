package follower

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entity struct {
    ID             primitive.ObjectID `bson:"_id,omitempty"`
    UserID         primitive.ObjectID `bson:"user_id"`
    FollowerUserID primitive.ObjectID `bson:"follower_user_id"`
}
