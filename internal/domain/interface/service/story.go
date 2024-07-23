package service

import "context"

type Story interface {
    Create(ctx context.Context, spec StoryCreateSpec) (result StoryCreateResult, err error)
    Delete(ctx context.Context, spec StoryDeleteSpec) (result StoryDeleteResult, err error)
}
