package story

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/util/timeutil"
)

type Service struct {
    time            timeutil.Time
    storyRepository repository.Story
}

func NewService(
    storyRepository repository.Story,
) service.Story {
    return &Service{
        time:            timeutil.NewTime(),
        storyRepository: storyRepository,
    }
}
