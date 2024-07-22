package session

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository) Create(ctx context.Context, spec repository.SessionCreateSpec) (result repository.SessionCreateResult, err error) {
    var entities []interface{}
    for _, model := range spec.Models {
        entities = append(entities, &Entity{
            UserID:    model.UserID,
            ExpiresAt: model.ExpiresAt,
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
