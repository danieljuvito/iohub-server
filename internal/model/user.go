package model

import (
    "errors"
    "github.com/danieljuvito/iohub-server/internal/util"
)

type User struct {
    ID       int
    Email    string
    Password string
}

func (u *User) Validate() error {
    if !util.ValidateEmail(u.Email) {
        return errors.New("invalid email")
    }

    if !util.ValidatePassword(u.Password) {
        return errors.New("invalid password")
    }

    return nil
}
