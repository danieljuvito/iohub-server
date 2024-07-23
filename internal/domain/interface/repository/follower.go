package repository

import "context"

type Follower interface {
    Get(ctx context.Context, spec FollowerGetSpec) (result FollowerGetResult, err error)

    Create(ctx context.Context, spec FollowerCreateSpec) (result FollowerCreateResult, err error)
    Delete(ctx context.Context, spec FollowerDeleteSpec) (result FollowerDeleteResult, err error)
}
