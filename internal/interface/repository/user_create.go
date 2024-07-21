package repository

import (
    "github.com/danieljuvito/iohub-server/internal/model"
)

type UserCreateSpec struct {
    *model.User
}

type UserCreateResult struct {
    ID int
}
