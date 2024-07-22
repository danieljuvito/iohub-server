package session

import (
    "context"
    "errors"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
)

func (s *Service) LogOut(ctx context.Context, spec service.SessionLogOutSpec) (result service.SessionLogOutResult, err error) {
    deleteResult, err := s.sessionRepository.Delete(ctx, repository.SessionDeleteSpec{
        IDs: []string{spec.SessionID},
    })
    if err != nil {
        return result, err
    }

    if deleteResult.Count == 0 {
        return result, errors.New("failed to delete")
    }

    return
}
