package repository

import "github.com/danieljuvito/iohub-server/internal/model"

type UserGetSpec struct {
    Email string
}

type UserGetResult struct {
    Data []*model.User
}
