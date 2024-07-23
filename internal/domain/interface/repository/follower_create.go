package repository

import "github.com/danieljuvito/iohub-server/internal/domain/model"

type FollowerCreateSpec struct {
    Data []*model.Follower
}

type FollowerCreateResult struct {
    IDs []string
}
