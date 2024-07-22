package repository

import (
    "github.com/danieljuvito/iohub-server/internal/domain/model"
)

type UserCreateSpec struct {
    Models []*model.User
}

type UserCreateResult struct {
    IDs []string
}
