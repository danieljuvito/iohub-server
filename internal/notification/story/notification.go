package story

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/notification"
    "golang.org/x/net/websocket"
)

type Notification struct {
    w *websocket.Conn
}

func NewNotification(w *websocket.Conn) notification.Story {
    return &Notification{
        w: w,
    }
}
