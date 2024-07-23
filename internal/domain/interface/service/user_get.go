package service

import (
    "github.com/danieljuvito/iohub-server/internal/domain/model"
)

type UserGetSpec struct {
    ID string
}

type UserGetResult struct {
    *model.User
}
