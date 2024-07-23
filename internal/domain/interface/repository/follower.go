package repository

import "context"

type Follower interface {
    Get(ctx context.Context, spec FollowerGetSpec) (result FollowerGetResult, err error)
}
