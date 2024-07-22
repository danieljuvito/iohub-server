package session

import (
    "github.com/danieljuvito/iohub-server/internal/interface/service"
    "github.com/labstack/echo/v4"
    "net/http"
)

type LogInRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LogInResponse struct {
    Token string `json:"token"`
}

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
            Token: res.Token,
        })
    })
}
