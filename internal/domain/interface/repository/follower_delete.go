package repository

import "github.com/danieljuvito/iohub-server/internal/domain/model"

type FollowerDeleteSpec struct {
    *model.Follower
}

type FollowerDeleteResult struct {
    Count int
}
