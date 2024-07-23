package user

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository) Get(ctx context.Context, spec repository.UserGetSpec) (result repository.UserGetResult, err error) {
    objectID, _ := primitive.ObjectIDFromHex(spec.ID)

    query := bson.D{
        {"$or", bson.A{
            bson.M{"_id": objectID},
            bson.M{"email": spec.Email},
        }},
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
        result.Data = append(result.Data, &model.User{
            ID:       entity.ID.Hex(),
            Name:     entity.Name,
            Email:    entity.Email,
            Password: entity.Password,
        })
    }

    return result, nil
}
