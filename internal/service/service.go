package service

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/repository"
    "github.com/danieljuvito/iohub-server/internal/service/session"
    "github.com/danieljuvito/iohub-server/internal/service/user"
)

type Service struct {
    UserService    service.User
    SessionService service.Session
}

func NewService(r *repository.Repository) *Service {
    return &Service{
        UserService:    user.NewService(r.UserRepository, r.FolloweeRepository),
        SessionService: session.NewService(r.UserRepository, r.SessionRepository),
    }
}
