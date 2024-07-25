package session

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/labstack/echo/v4"
    "net/http"
)

type LogInRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LogInResponse struct {
    UserID string `json:"user_id"`
    Token  string `json:"token"`
}

// LogIn handles user login requests.
//
// @Summary User login
// @Description Authenticates a user based on email and password.
// @Tags Sessions
// @Accept json
// @Produce json
// @Param request body LogInRequest true "User login request"
// @Success 200 {object} LogInResponse "Successful login"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Router /log-in [post]
func (c *controller) LogIn() {
    c.POST("/log-in", func(e echo.Context) error {
        ctx := e.Request().Context()

        req := new(LogInRequest)
        if err := e.Bind(req); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        }

        res, err := c.sessionService.LogIn(ctx, service.SessionLogInSpec{
            Email:    req.Email,
            Password: req.Password,
        })
        if err != nil {
            return err
        }

        return e.JSON(http.StatusOK, LogInResponse{
            UserID: res.UserID,
            Token:  res.Token,
        })
    })
}
