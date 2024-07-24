package repository

import "github.com/danieljuvito/iohub-server/internal/domain/model"

type UserGetSpec struct {
    IDs   []string
    Email string
    Name  string
    Page  int
    Limit int
}

type UserGetResult struct {
    Data []*model.User
}
