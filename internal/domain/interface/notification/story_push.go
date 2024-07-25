package notification

type StoryPushSpec struct {
    FolloweeID  string   `json:"followee_id"`
    FollowerIDs []string `json:"follower_ids"`
}

type StoryPushResult struct {
}
