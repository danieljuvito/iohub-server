package repository

import "context"

type Session interface {
    Create(ctx context.Context, spec SessionCreateSpec) (result SessionCreateResult, err error)
    Delete(ctx context.Context, spec SessionDeleteSpec) (result SessionDeleteResult, err error)
}
