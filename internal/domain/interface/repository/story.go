package repository

import "context"

type Story interface {
    Get(ctx context.Context, specs ...StoryGetSpec) (result StoryGetResult, err error)

    Create(ctx context.Context, spec StoryCreateSpec) (result StoryCreateResult, err error)
    Delete(ctx context.Context, spec StoryDeleteSpec) (result StoryDeleteResult, err error)
}
