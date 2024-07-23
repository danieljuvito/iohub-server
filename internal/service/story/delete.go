package story

import (
    "context"
    "errors"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
)

func (s *Service) Delete(ctx context.Context, spec service.StoryDeleteSpec) (result service.StoryDeleteResult, err error) {
    deleteResult, err := s.storyRepository.Delete(ctx, repository.StoryDeleteSpec{
        IDs: []string{spec.ID},
    })
    if err != nil {
        return result, err
    }

    if deleteResult.Count == 0 {
        return result, errors.New("failed to delete")
    }

    result.Count = deleteResult.Count

    return result, nil
}
