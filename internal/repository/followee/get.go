package followee

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository) Get(ctx context.Context, spec repository.FolloweeGetSpec) (result repository.FolloweeGetResult, err error) {
    orQuery := bson.A{}
    if len(spec.UserID) != 0 {
        objectID, _ := primitive.ObjectIDFromHex(spec.UserID)
        orQuery = append(orQuery, bson.M{"user_id": objectID})
    }
    query := bson.D{}
    if len(orQuery) != 0 {
        query = append(query, bson.E{Key: "$or", Value: orQuery})
    }
    res, err := r.collection.Find(ctx, query)
    if err != nil {
        return result, err
    }

    var entities []*Entity
    err = res.All(ctx, &entities)
    if err != nil {
        return result, err
    }

    for _, entity := range entities {
        result.Data = append(result.Data, &model.Followee{
            ID:             entity.ID.Hex(),
            UserID:         entity.UserID.Hex(),
            FolloweeUserID: entity.FolloweeUserID.Hex(),
        })
    }

    return result, nil
}
