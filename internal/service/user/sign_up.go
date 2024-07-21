package user

import (
    "context"
    "github.com/danieljuvito/iohub-server/internal/interface/repository"
    "github.com/danieljuvito/iohub-server/internal/interface/service"
    "github.com/danieljuvito/iohub-server/internal/util/errorutil"
    "golang.org/x/crypto/bcrypt"
)

func (s *Service) SignUp(ctx context.Context, spec service.UserSignUpSpec) (result service.UserSignUpResult, err error) {
    err = spec.Validate()
    if err != nil {
        return result, err
    }

    getResult, err := s.userRepository.Get(ctx, repository.UserGetSpec{
        Email: spec.Email,
    })
    if err != nil {
        return result, err
    }
    if len(getResult.Data) != 0 {
        return result, errorutil.AlreadyExistError.Wrap("user is already exist")
    }

    // Generate a salted hash
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(spec.Password), bcrypt.DefaultCost)
    if err != nil {
        return result, err
    }
    spec.User.Password = string(hashedPassword)

    createResult, err := s.userRepository.Create(ctx, repository.UserCreateSpec{
        User: spec.User,
    })
    if err != nil {
        return result, err
    }

    return service.UserSignUpResult{
        ID: createResult.ID,
    }, nil
}
