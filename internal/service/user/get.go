package user

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/util/errorutil"
)

func (s *Service) Get(ctx context.Context, spec service.UserGetSpec) (result service.UserGetResult, err error) {
    getResult, err := s.userRepository.Get(ctx, repository.UserGetSpec{
        IDs: []string{spec.ID},
    })
    if err != nil {
        return result, err
    }

    if len(getResult.Data) != 1 {
        return result, errorutil.NotFound.Wrap("user not found")
    }

    result.User = getResult.Data[0]

    return result, nil
}
