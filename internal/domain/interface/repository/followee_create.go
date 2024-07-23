package repository

import "github.com/danieljuvito/iohub-server/internal/domain/model"

type FolloweeCreateSpec struct {
    Data []*model.Followee
}

type FolloweeCreateResult struct {
    IDs []string
}
