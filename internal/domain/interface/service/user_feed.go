package service

import (
    "github.com/danieljuvito/iohub-server/internal/domain/model"
)

type UserFeedSpec struct {
    UserID string
    Page   int
    Limit  int
}

type UserFeedResult struct {
    Data []UserFeedResultData
}

type UserFeedResultData struct {
    *model.User
    Stories []*model.Story
}
