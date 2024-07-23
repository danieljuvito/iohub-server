package service

import "context"

type User interface {
    Get(ctx context.Context, spec UserGetSpec) (result UserGetResult, err error)
    List(ctx context.Context, spec UserListSpec) (result UserListResult, err error)

    SignUp(context.Context, UserSignUpSpec) (UserSignUpResult, error)
}
