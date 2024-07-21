package model

import (
    "github.com/danieljuvito/iohub-server/internal/util/errorutil"
    "github.com/danieljuvito/iohub-server/internal/util/validate"
)

type User struct {
    ID       int
    Email    string
    Password string
}

func (u *User) Validate() error {
    if !validate.Email(u.Email) {
        return errorutil.InvalidError.Wrap("invalid email")
    }

    if !validate.Password(u.Password) {
        return errorutil.InvalidError.Wrap("invalid password")
    }

    return nil
}
