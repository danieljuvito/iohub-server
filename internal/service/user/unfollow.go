package user

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "github.com/danieljuvito/iohub-server/internal/util/errorutil"
)

func (s *Service) Unfollow(ctx context.Context, spec service.UserUnfollowSpec) (result service.UserUnfollowResult, err error) {
    getResult, err := s.userRepository.Get(ctx, repository.UserGetSpec{
        ID: spec.FolloweeUserID,
    })
    if err != nil {
        return result, err
    }
    if len(getResult.Data) != 1 {
        return result, errorutil.NotFound.Wrap("user not found")
    }

    _, err = s.followeeRepository.Create(ctx, repository.FolloweeCreateSpec{
        Data: []*model.Followee{
            {
                UserID:         spec.UserID,
                FolloweeUserID: spec.FolloweeUserID,
            },
        },
    })
    if err != nil {
        return result, err
    }

    _, err = s.followerRepository.Create(ctx, repository.FollowerCreateSpec{
        Data: []*model.Follower{
            {
                UserID:         spec.FolloweeUserID,
                FollowerUserID: spec.UserID,
            },
        },
    })
    if err != nil {
        return result, err
    }

    return result, nil
}
