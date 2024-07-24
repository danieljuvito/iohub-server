package repository

import (
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "time"
)

type UserGetFolloweeStoriesSpec struct {
    UserID       string
    ExpiresAfter time.Time
    Page         int
    Limit        int
}

type UserGetFolloweeStoriesResult struct {
    Data []*model.User
}
