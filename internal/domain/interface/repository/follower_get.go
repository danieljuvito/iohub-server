package repository

import "github.com/danieljuvito/iohub-server/internal/domain/model"

type FollowerGetSpec struct {
    UserID string
}

type FollowerGetResult struct {
    Data []*model.Follower
}
