package repository

import (
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "time"
)

type FolloweeGetSpec struct {
    UserID    string
    WithStory bool
    ExpiresAt time.Time
    Limit     int
    Page      int
}

type FolloweeGetResult struct {
    Data []*model.Followee
}
