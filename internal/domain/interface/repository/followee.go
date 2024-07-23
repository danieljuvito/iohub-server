package repository

import "context"

type Followee interface {
    Get(ctx context.Context, spec FolloweeGetSpec) (result FolloweeGetResult, err error)

    Create(ctx context.Context, spec FolloweeCreateSpec) (result FolloweeCreateResult, err error)
    Delete(ctx context.Context, spec FolloweeDeleteSpec) (result FolloweeDeleteResult, err error)
}
