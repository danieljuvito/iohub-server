package user

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
)

func (s *Service) List(ctx context.Context, spec service.UserListSpec) (result service.UserListResult, err error) {
    getUserResult, err := s.userRepository.Get(ctx, repository.UserGetSpec{
        Name:  spec.Name,
        Limit: spec.Limit,
        Page:  spec.Page,
    })
    if err != nil {
        return result, err
    }

    getFolloweeResult, err := s.followeeRepository.Get(ctx, repository.FolloweeGetSpec{
        UserID: spec.UserID,
    })
    if err != nil {
        return result, err
    }

    followeeByFolloweeUserID := make(map[string]*model.Followee)
    for _, followee := range getFolloweeResult.Data {
        followeeByFolloweeUserID[followee.FolloweeUserID] = followee
    }

    for _, user := range getUserResult.Data {
        result.Data = append(result.Data, service.UserListResultData{
            User:     user,
            Followee: followeeByFolloweeUserID[user.ID],
        })
    }

    return result, nil
}
