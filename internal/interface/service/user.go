package service

import "context"

type User interface {
    SignUp(context.Context, UserSignUpSpec) (UserSignUpResult, error)
}
