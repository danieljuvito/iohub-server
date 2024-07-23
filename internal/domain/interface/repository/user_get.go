package repository

import "github.com/danieljuvito/iohub-server/internal/domain/model"

type UserGetSpec struct {
    ID    string
    Email string
}

type UserGetResult struct {
    Data []*model.User
}
