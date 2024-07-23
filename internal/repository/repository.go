package repository

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/repository/followee"
    "github.com/danieljuvito/iohub-server/internal/repository/follower"
    "github.com/danieljuvito/iohub-server/internal/repository/session"
    "github.com/danieljuvito/iohub-server/internal/repository/story"
    "github.com/danieljuvito/iohub-server/internal/repository/user"
    "go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
    UserRepository     repository.User
    SessionRepository  repository.Session
    FolloweeRepository repository.Followee
    FollowerRepository repository.Follower
    StoryRepository    repository.Story
}

func NewRepository(db *mongo.Database) *Repository {
    return &Repository{
        UserRepository:     user.NewRepository(db),
        SessionRepository:  session.NewRepository(db),
        FolloweeRepository: followee.NewRepository(db),
        FollowerRepository: follower.NewRepository(db),
        StoryRepository:    story.NewRepository(db),
    }
}
