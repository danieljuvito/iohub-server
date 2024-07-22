package service

import "context"

type Session interface {
    LogIn(ctx context.Context, spec SessionLogInSpec) (result SessionLogInResult, err error)
    LogOut(ctx context.Context, spec SessionLogOutSpec) (result SessionLogOutResult, err error)
}
