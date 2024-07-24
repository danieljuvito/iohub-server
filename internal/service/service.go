package service

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/repository"
    "github.com/danieljuvito/iohub-server/internal/service/session"
    "github.com/danieljuvito/iohub-server/internal/service/story"
    "github.com/danieljuvito/iohub-server/internal/service/user"
)

type Service struct {
    UserService    service.User
    SessionService service.Session
    StoryService   service.Story
}

func NewService(r *repository.Repository) *Service {
    return &Service{
        UserService:    user.NewService(r.UserRepository, r.FolloweeRepository, r.FollowerRepository, r.StoryRepository),
        SessionService: session.NewService(r.UserRepository, r.SessionRepository),
        StoryService:   story.NewService(r.StoryRepository),
    }
}
