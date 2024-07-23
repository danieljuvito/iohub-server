package follower

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) Delete(ctx context.Context, spec repository.FollowerDeleteSpec) (result repository.FollowerDeleteResult, err error) {
    userID, _ := primitive.ObjectIDFromHex(spec.UserID)
    followerUserID, _ := primitive.ObjectIDFromHex(spec.FollowerUserID)

    res, err := r.collection.DeleteMany(ctx, bson.M{
        "user_id":          userID,
        "follower_user_id": followerUserID,
    })
    if err != nil {
        return result, err
    }

    result.Count = int(res.DeletedCount)

    return result, nil
}
