package timeutil

import "time"

type Time interface {
    Now() time.Time
}

type timeutil struct {
}

func (t *timeutil) Now() time.Time {
    return time.Now()
}

func NewTime() Time {
    return &timeutil{}
}
