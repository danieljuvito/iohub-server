package story

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/notification"
    "golang.org/x/net/websocket"
)

type Notification struct {
    ws *websocket.Conn
}

func NewNotification(ws *websocket.Conn) notification.Story {
    return &Notification{
        ws: ws,
    }
}
