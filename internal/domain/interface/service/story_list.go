package service

import (
    "github.com/danieljuvito/iohub-server/internal/domain/model"
)

type StoryGetSpec struct {
    ID string
}

type StoryGetResult struct {
    *model.Story
}
