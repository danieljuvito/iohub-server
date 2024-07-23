package model

import (
    "github.com/danieljuvito/iohub-server/internal/util/errorutil"
    "github.com/danieljuvito/iohub-server/internal/util/validate"
)

type User struct {
    ID       string
    Name     string
    Email    string
    Password string
}

func (u *User) Validate() error {
    if !validate.Email(u.Email) {
        return errorutil.Invalid.Wrap("invalid email")
    }

    if !validate.Password(u.Password) {
        return errorutil.Invalid.Wrap("invalid password")
    }

    return nil
}
