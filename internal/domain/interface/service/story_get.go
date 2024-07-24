package service

import (
    "github.com/danieljuvito/iohub-server/internal/domain/model"
)

type StoryListSpec struct {
    UserID string
}

type StoryListResult struct {
    Data []*model.Story
}
