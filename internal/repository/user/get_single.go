package user

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/interface/repository"
    "go.mongodb.org/mongo-driver/bson"
)

func (r Repository) Get(ctx context.Context, spec repository.UserGetSpec) (result repository.UserGetResult, err error) {
    cursor, err := r.db.Collection("users").Find(ctx, bson.D{{"email", spec.Email}})
    if err != nil {
        return result, err
    }

    err = cursor.All(ctx, &result.Data)
    if err != nil {
        return result, err
    }

    return result, nil
}
