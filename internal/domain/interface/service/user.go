package service

import "context"

type User interface {
    Get(ctx context.Context, spec UserGetSpec) (result UserGetResult, err error)
    List(ctx context.Context, spec UserListSpec) (result UserListResult, err error)
    Feed(ctx context.Context, spec UserFeedSpec) (result UserFeedResult, err error)

    SignUp(ctx context.Context, spec UserSignUpSpec) (result UserSignUpResult, err error)
    Follow(ctx context.Context, spec UserFollowSpec) (result UserFollowResult, err error)
    Unfollow(ctx context.Context, spec UserUnfollowSpec) (result UserUnfollowResult, err error)
}
