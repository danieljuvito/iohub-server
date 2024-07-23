package followee

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) Delete(ctx context.Context, spec repository.FolloweeDeleteSpec) (result repository.FolloweeDeleteResult, err error) {
    userID, _ := primitive.ObjectIDFromHex(spec.UserID)
    followeeUserID, _ := primitive.ObjectIDFromHex(spec.FolloweeUserID)

    res, err := r.collection.DeleteOne(ctx, bson.M{
        "user_id":          userID,
        "followee_user_id": followeeUserID,
    })
    if err != nil {
        return result, err
    }

    result.Count = int(res.DeletedCount)

    return result, nil
}
