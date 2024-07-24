package user

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/util/timeutil"
)

type Service struct {
    time               timeutil.Time
    userRepository     repository.User
    followeeRepository repository.Followee
    followerRepository repository.Follower
    storyRepository    repository.Story
}

func NewService(
    userRepository repository.User,
    followeeRepository repository.Followee,
    followerRepository repository.Follower,
    storyRepository repository.Story,
) service.User {
    return &Service{
        time:               timeutil.NewTime(),
        userRepository:     userRepository,
        followeeRepository: followeeRepository,
        followerRepository: followerRepository,
        storyRepository:    storyRepository,
    }
}
