package story

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
)

func (s *Service) List(ctx context.Context, spec service.StoryListSpec) (result service.StoryListResult, err error) {
    getStoryResult, err := s.storyRepository.Get(ctx,
        repository.StoryGetSpec{
            UserIDs: []string{spec.UserID},
        },
        repository.StoryGetSpec{
            ExpiresAfter: s.time.Now(),
        },
    )
    if err != nil {
        return result, err
    }

    result.Data = getStoryResult.Data

    return result, nil
}
