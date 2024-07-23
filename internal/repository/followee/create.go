package followee

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) Create(ctx context.Context, spec repository.FolloweeCreateSpec) (result repository.FolloweeCreateResult, err error) {
    var entities []interface{}
    for _, model := range spec.Data {
        userID, err := primitive.ObjectIDFromHex(model.UserID)
        if err != nil {
            return result, nil
        }
        followerUserID, err := primitive.ObjectIDFromHex(model.FolloweeUserID)
        if err != nil {
            return result, nil
        }
        entities = append(entities, &Entity{
            UserID:         userID,
            FolloweeUserID: followerUserID,
        })
    }

    res, err := r.collection.InsertMany(ctx, entities)
    if err != nil {
        return result, err
    }

    for _, id := range res.InsertedIDs {
        result.IDs = append(result.IDs, (id).(primitive.ObjectID).Hex())
    }

    return result, nil
}
