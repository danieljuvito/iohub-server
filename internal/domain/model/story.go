package model

import "time"

type Story struct {
    ID        string
    UserID    string
    Content   string
    ExpiresAt time.Time
}
