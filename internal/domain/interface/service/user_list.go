package service

import (
    "github.com/danieljuvito/iohub-server/internal/domain/model"
)

type UserListSpec struct {
    UserID string
    Name   string
    Page   int
    Limit  int
}

type UserListResult struct {
    Data []UserListResultData
}

type UserListResultData struct {
    *model.User
    *model.Followee
}
