package service

import (
    "github.com/danieljuvito/iohub-server/internal/domain/model"
)

type UserSignUpSpec struct {
    *model.User
}

type UserSignUpResult struct {
    ID string
}
