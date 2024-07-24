package notification

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/notification"
    "github.com/danieljuvito/iohub-server/internal/notification/story"
    "golang.org/x/net/websocket"
)

type Notification struct {
    StoryNotification notification.Story
}

func NewNotification(w *websocket.Conn) *Notification {
    return &Notification{
        StoryNotification: story.NewNotification(w),
    }
}
