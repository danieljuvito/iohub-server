package session

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
)

func (s *Service) Create(ctx context.Context, spec service.StoryCreateSpec) (result service.StoryCreateResult, err error) {
    createResult, err := s.storyRepository.Create(ctx, repository.StoryCreateSpec{
        Models: []*model.Story{
            spec.Story,
        },
    })
    if err != nil {
        return result, err
    }

    result.ID = createResult.IDs[0]

    return result, nil
}
