package repository

import "github.com/danieljuvito/iohub-server/internal/domain/model"

type SessionCreateSpec struct {
    Models []*model.Session
}

type SessionCreateResult struct {
    IDs []string
}
