package repository

type FolloweeDeleteSpec struct {
    IDs []string
}

type FolloweeDeleteResult struct {
    Count int
}
