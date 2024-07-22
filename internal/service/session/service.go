package session

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
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
