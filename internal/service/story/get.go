package story

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/util/errorutil"
)

func (s *Service) Get(ctx context.Context, spec service.StoryGetSpec) (result service.StoryGetResult, err error) {
    getStoryResult, err := s.storyRepository.Get(ctx,
        repository.StoryGetSpec{
            IDs: []string{spec.ID},
        },
    )
    if err != nil {
        return result, err
    }

    if len(getStoryResult.Data) != 1 {
        return result, errorutil.NotFound
    }

    result.Story = getStoryResult.Data[0]

    return result, nil
}
