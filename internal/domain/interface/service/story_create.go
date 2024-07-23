package service

import "github.com/danieljuvito/iohub-server/internal/domain/model"

type StoryCreateSpec struct {
    *model.Story
}

type StoryCreateResult struct {
    ID string
}
