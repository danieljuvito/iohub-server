package session

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
)

type Service struct {
    userRepository  repository.User
    storyRepository repository.Story
}

func NewService(
    userRepository repository.User,
    storyRepository repository.Story,
) service.Story {
    return &Service{
        userRepository:  userRepository,
        storyRepository: storyRepository,
    }
}
