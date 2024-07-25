package story

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/notification"
    "github.com/danieljuvito/iohub-server/internal/util/websocketutil"
)

type Notification struct {
    hub *websocketutil.Hub
}

func NewNotification(hub *websocketutil.Hub) notification.Story {
    return &Notification{
        hub: hub,
    }
}
