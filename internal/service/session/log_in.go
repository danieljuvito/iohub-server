package session

import (
    "context"
    //    "encoding/base64"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "github.com/danieljuvito/iohub-server/internal/util/errorutil"
    "github.com/danieljuvito/iohub-server/internal/util/jwtutil"
    "golang.org/x/crypto/bcrypt"
    "time"
)

func (s *Service) LogIn(ctx context.Context, spec service.SessionLogInSpec) (result service.SessionLogInResult, err error) {
    getResult, err := s.userRepository.Get(ctx, repository.UserGetSpec{
        Email: spec.Email,
    })
    if err != nil {
        return result, err
    }
    if len(getResult.Data) != 1 {
        return result, errorutil.Unauthorized
    }

    user := getResult.Data[0]

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(spec.Password))
    if err != nil {
        return result, errorutil.Unauthorized
    }

    session := model.Session{
        UserID:    user.ID,
        ExpiresAt: s.time.Now().Add(time.Hour * 24),
    }
    createResult, err := s.sessionRepository.Create(ctx, repository.SessionCreateSpec{
        Models: []*model.Session{
            &session,
        },
    })
    if err != nil {
        return result, err
    }

    token, err := jwtutil.GenerateToken(createResult.IDs[0], session.ExpiresAt)
    if err != nil {
        return result, err
    }

    result.Token = token

    return result, nil
}
