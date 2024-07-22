package user

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
)

type Service struct {
    userRepository repository.User
}

func NewService(userRepository repository.User) service.User {
    return &Service{
        userRepository: userRepository,
    }
}
