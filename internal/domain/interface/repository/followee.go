package repository

import "context"

type Followee interface {
    Get(ctx context.Context, spec FolloweeGetSpec) (result FolloweeGetResult, err error)
}
