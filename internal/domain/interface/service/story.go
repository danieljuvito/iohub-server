package service

import "context"

type Story interface {
    Get(ctx context.Context, spec StoryGetSpec) (result StoryGetResult, err error)
    List(ctx context.Context, spec StoryListSpec) (result StoryListResult, err error)

    Create(ctx context.Context, spec StoryCreateSpec) (result StoryCreateResult, err error)
    Delete(ctx context.Context, spec StoryDeleteSpec) (result StoryDeleteResult, err error)
}
