package user

import (
    "github.com/danieljuvito/iohub-server/internal/interface/repository"
    "go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
    db *mongo.Database
}

func NewRepository(db *mongo.Database) repository.User {
    return &Repository{
        db: db,
    }
}
