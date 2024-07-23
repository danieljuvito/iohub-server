package followee

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entity struct {
    ID             primitive.ObjectID `bson:"_id,omitempty"`
    UserID         primitive.ObjectID `bson:"user_id"`
    FolloweeUserID primitive.ObjectID `bson:"followee_user_id"`
}
