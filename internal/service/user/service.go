package user

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
)

type Service struct {
    userRepository     repository.User
    followeeRepository repository.Followee
}

func NewService(
    userRepository repository.User,
    followeeRepository repository.Followee,
) service.User {
    return &Service{
        userRepository:     userRepository,
        followeeRepository: followeeRepository,
    }
}
