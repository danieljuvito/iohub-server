package story

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
    collection *mongo.Collection
}

func NewRepository(db *mongo.Database) repository.Story {
    return &Repository{
        collection: db.Collection("stories"),
    }
}
