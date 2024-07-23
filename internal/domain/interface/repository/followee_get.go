package repository

import "github.com/danieljuvito/iohub-server/internal/domain/model"

type FolloweeGetSpec struct {
    UserID string
}

type FolloweeGetResult struct {
    Data []*model.Followee
}
