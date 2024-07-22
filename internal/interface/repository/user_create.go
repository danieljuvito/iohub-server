package repository

import (
    "github.com/danieljuvito/iohub-server/internal/model"
)

type UserCreateSpec struct {
    Models []*model.User
}

type UserCreateResult struct {
    IDs []string
}
