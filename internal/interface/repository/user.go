package repository

import "context"

type User interface {
    Get(ctx context.Context, spec UserGetSpec) (result UserGetResult, err error)

    Create(ctx context.Context, spec UserCreateSpec) (result UserCreateResult, err error)
}
