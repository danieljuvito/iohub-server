package user

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/model"
    "go.mongodb.org/mongo-driver/bson"
)

func (r Repository) Get(ctx context.Context, spec repository.UserGetSpec) (result repository.UserGetResult, err error) {
    res, err := r.collection.Find(ctx, bson.D{{"email", spec.Email}})
    if err != nil {
        return result, err
    }

    var entities []*Entity
    err = res.All(ctx, &entities)
    if err != nil {
        return result, err
    }

    for _, entitiy := range entities {
        result.Data = append(result.Data, &model.User{
            ID:       entitiy.ID.Hex(),
            Email:    entitiy.Email,
            Password: entitiy.Password,
        })
    }

    return result, nil
}
