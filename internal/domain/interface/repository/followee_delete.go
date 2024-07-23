package repository

import "github.com/danieljuvito/iohub-server/internal/domain/model"

type FolloweeDeleteSpec struct {
    *model.Followee
}

type FolloweeDeleteResult struct {
    Count int
}
