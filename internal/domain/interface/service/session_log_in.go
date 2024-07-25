package service

type SessionLogInSpec struct {
    Email    string
    Password string
}

type SessionLogInResult struct {
    UserID string
    Token  string
}
