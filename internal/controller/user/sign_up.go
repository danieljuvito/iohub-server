package user

import (
    "errors"
    "github.com/danieljuvito/iohub-server/internal/interface/service"
    "github.com/danieljuvito/iohub-server/internal/model"
    "github.com/danieljuvito/iohub-server/internal/util/errorutil"
    "github.com/labstack/echo/v4"
    "net/http"
)

type SignUpRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (c *controller) SignUp() {
    c.POST("/sign-up", func(e echo.Context) error {
        ctx := e.Request().Context()

        request := new(SignUpRequest)
        if err := e.Bind(request); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        }

        _, err := c.userService.SignUp(ctx, service.UserSignUpSpec{
            User: &model.User{
                Email:    request.Email,
                Password: request.Password,
            },
        })
        if err != nil {
            if errors.Is(err, errorutil.InvalidError) {
                return echo.NewHTTPError(http.StatusBadRequest, err.Error())
            }
            if errors.Is(err, errorutil.AlreadyExistError) {
                return echo.NewHTTPError(http.StatusConflict, err.Error())
            }
            return err
        }
        return e.JSON(http.StatusOK, "OK")
    })
}
