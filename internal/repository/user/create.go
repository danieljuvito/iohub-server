package user

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/interface/repository"
)

func (r Repository) Create(ctx context.Context, spec repository.UserCreateSpec) (result repository.UserCreateResult, err error) {
    _, err = r.db.Collection("users").InsertOne(ctx, spec.User)
    if err != nil {
        return result, err
    }

    return
}
