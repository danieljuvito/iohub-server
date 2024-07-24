package user

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
)

func (s *Service) Feed(ctx context.Context, spec service.UserFeedSpec) (result service.UserFeedResult, err error) {
    getFolloweeResult, err := s.followeeRepository.Get(ctx, repository.FolloweeGetSpec{
        UserID:    spec.UserID,
        WithStory: true,
        Limit:     spec.Limit,
        Page:      spec.Page,
        ExpiresAt: s.time.Now(),
    })
    if err != nil {
        return result, err
    }

    followeeUserIDs := make([]string, 0)
    for _, followee := range getFolloweeResult.Data {
        followeeUserIDs = append(followeeUserIDs, followee.FolloweeUserID)
    }

    getStoryResult, err := s.storyRepository.Get(ctx, repository.StoryGetSpec{
        UserIDs:      followeeUserIDs,
        ExpiresAfter: s.time.Now(),
    })
    if err != nil {
        return result, err
    }

    storiesByUserID := make(map[string][]*model.Story)
    for _, story := range getStoryResult.Data {
        storiesByUserID[story.UserID] = append(storiesByUserID[story.UserID], story)
    }

    getUserResult, err := s.userRepository.Get(ctx, repository.UserGetSpec{
        IDs: followeeUserIDs,
    })
    if err != nil {
        return result, err
    }

    for _, user := range getUserResult.Data {
        result.Data = append(result.Data, service.UserFeedResultData{
            User:    user,
            Stories: storiesByUserID[user.ID],
        })
    }

    return result, nil
}
