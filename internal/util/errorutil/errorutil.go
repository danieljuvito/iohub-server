package errorutil

import (
    "errors"
    "fmt"
)

type Error struct {
    error
}

var (
    Invalid      = Error{errors.New("invalid")}
    AlreadyExist = Error{errors.New("already exist")}
    NotFound     = Error{errors.New("not found")}
    Unauthorized = Error{errors.New("unauthorized")}
)

func (err Error) Wrap(msg string) error {
    if err.error == nil {
        return errors.New(msg)
    }
    return fmt.Errorf("%s: %w", msg, err)
}
