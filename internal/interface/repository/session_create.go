package repository

import "github.com/danieljuvito/iohub-server/internal/model"

type SessionCreateSpec struct {
    Models []*model.Session
}

type SessionCreateResult struct {
    IDs []string
}
