package errorutil

import (
    "errors"
    "fmt"
)

type Error struct {
    error
}

var (
    InvalidError      = Error{errors.New("invalid")}
    AlreadyExistError = Error{errors.New("already exist")}
)

func (err Error) Wrap(msg string) error {
    if err.error == nil {
        return errors.New(msg)
    }
    return fmt.Errorf("%s: %w", msg, err)
}
