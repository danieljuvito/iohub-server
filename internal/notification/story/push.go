package story

import (
    "context"
    "encoding/json"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/notification"
)

func (n *Notification) Push(ctx context.Context, spec notification.StoryPushSpec) (result notification.StoryPushResult, err error) {
    marshal, err := json.Marshal(spec)
    if err != nil {
        return result, err
    }
    n.hub.Send(marshal)
    if err != nil {
        return result, err
    }

    return result, nil
}
