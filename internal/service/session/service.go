package session

import (
    "github.com/danieljuvito/iohub-server/internal/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/interface/service"
)

type Service struct {
    userRepository    repository.User
    sessionRepository repository.Session
}

func NewService(
    userRepository repository.User,
    sessionRepository repository.Session,
) service.Session {
    return &Service{
        userRepository:    userRepository,
        sessionRepository: sessionRepository,
    }
}
