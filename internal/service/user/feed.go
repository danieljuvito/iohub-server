package user

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
)

func (s *Service) Feed(ctx context.Context, spec service.UserFeedSpec) (result service.UserFeedResult, err error) {
    getFolloweeStoriesResult, err := s.userRepository.GetFolloweeStories(ctx, repository.UserGetFolloweeStoriesSpec{
        UserID:       spec.UserID,
        ExpiresAfter: s.time.Now(),
        Limit:        spec.Limit,
        Page:         spec.Page,
    })
    if err != nil {
        return result, err
    }

    followeeUserIDs := make([]string, 0)
    for _, followee := range getFolloweeStoriesResult.Data {
        followeeUserIDs = append(followeeUserIDs, followee.ID)
    }

    getStoryResult, err := s.storyRepository.Get(ctx, repository.StoryGetSpec{
        UserIDs: followeeUserIDs,
    })
    if err != nil {
        return result, err
    }

    storiesByUserID := make(map[string][]*model.Story)
    for _, story := range getStoryResult.Data {
        storiesByUserID[story.UserID] = append(storiesByUserID[story.UserID], story)
    }

    for _, user := range getFolloweeStoriesResult.Data {
        result.Data = append(result.Data, service.UserFeedResultData{
            User:    user,
            Stories: storiesByUserID[user.ID],
        })
    }

    return result, nil
}
