package session

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/util/timeutil"
)

type Service struct {
    time              timeutil.Time
    userRepository    repository.User
    sessionRepository repository.Session
}

func NewService(
    userRepository repository.User,
    sessionRepository repository.Session,
) service.Session {
    return &Service{
        time:              timeutil.NewTime(),
        userRepository:    userRepository,
        sessionRepository: sessionRepository,
    }
}
