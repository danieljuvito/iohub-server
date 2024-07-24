package notification

import "context"

type Story interface {
    Push(ctx context.Context, spec StoryPushSpec) (result StoryPushResult, err error)
}
