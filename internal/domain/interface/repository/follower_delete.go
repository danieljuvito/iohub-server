package repository

type FollowerDeleteSpec struct {
    IDs []string
}

type FollowerDeleteResult struct {
    Count int
}
