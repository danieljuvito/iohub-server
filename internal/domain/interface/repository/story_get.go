package repository

import (
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "time"
)

type StoryGetSpec struct {
    IDs          []string
    UserIDs      []string
    ExpiresAfter time.Time
}

type StoryGetResult struct {
    Data []*model.Story
}
