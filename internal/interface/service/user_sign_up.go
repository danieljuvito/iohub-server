package service

import (
    "github.com/danieljuvito/iohub-server/internal/model"
)

type UserSignUpSpec struct {
    *model.User
}

type UserSignUpResult struct {
    ID int
}
