package user

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository) Create(ctx context.Context, spec repository.UserCreateSpec) (result repository.UserCreateResult, err error) {
    var entities []interface{}
    for _, model := range spec.Models {
        entities = append(entities, &Entity{
            Email:    model.Email,
            Password: model.Password,
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
