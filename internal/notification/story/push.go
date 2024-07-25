package story

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/notification"
    "golang.org/x/net/websocket"
)

func (n *Notification) Push(ctx context.Context, spec notification.StoryPushSpec) (result notification.StoryPushResult, err error) {
    err = websocket.JSON.Send(n.ws, spec)
    if err != nil {
        return result, err
    }

    return result, nil
}
