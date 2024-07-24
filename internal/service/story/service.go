package story

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/notification"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/util/timeutil"
)

type Service struct {
    time               timeutil.Time
    storyRepository    repository.Story
    followerRepository repository.Follower
    storyNotification  notification.Story
}

func NewService(
    storyRepository repository.Story,
    followerRepository repository.Follower,
    storyNotification notification.Story,
) service.Story {
    return &Service{
        time:               timeutil.NewTime(),
        storyRepository:    storyRepository,
        followerRepository: followerRepository,
        storyNotification:  storyNotification,
    }
}
