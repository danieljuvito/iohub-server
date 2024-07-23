package repository

import "github.com/danieljuvito/iohub-server/internal/domain/model"

type StoryCreateSpec struct {
    Models []*model.Story
}

type StoryCreateResult struct {
    IDs []string
}
