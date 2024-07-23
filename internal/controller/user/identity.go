package user

import (
    "github.com/danieljuvito/iohub-server/internal/controller/middleware"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "github.com/labstack/echo/v4"
    "net/http"
)

type IdentityResponse struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func (c *controller) Identity() {
    c.GET("/identity", func(e echo.Context) error {
        ctx := e.Request().Context()

        var session = e.Get("session").(model.Session)

        res, err := c.userService.Get(ctx, service.UserGetSpec{
            ID: session.UserID,
        })
        if err != nil {
            return err
        }
        return e.JSON(http.StatusOK, IdentityResponse{
            ID:    res.ID,
            Name:  res.Name,
            Email: res.Email,
        })
    }, middleware.Auth)
}
