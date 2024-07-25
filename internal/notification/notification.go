package notification

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/notification"
    "github.com/danieljuvito/iohub-server/internal/notification/story"
    "github.com/danieljuvito/iohub-server/internal/util/websocketutil"
)

type Notification struct {
    StoryNotification notification.Story
}

func NewNotification(hub *websocketutil.Hub) *Notification {
    return &Notification{
        StoryNotification: story.NewNotification(hub),
    }
}
