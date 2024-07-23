package service

import (
    "github.com/danieljuvito/iohub-server/internal/domain/model"
)

type UserListSpec struct {
    Search string
}

type UserListResult struct {
    Data []*model.User
}
