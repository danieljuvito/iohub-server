package repository

type SessionDeleteSpec struct {
    IDs []string
}

type SessionDeleteResult struct {
    Count int
}
