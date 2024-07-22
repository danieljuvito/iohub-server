package service

type SessionLogInSpec struct {
    Email    string
    Password string
}

type SessionLogInResult struct {
    Token string
}
