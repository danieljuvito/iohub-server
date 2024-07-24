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

// Identity handles GET requests for retrieving user identity details.
// @Summary Get user identity
// @Description Retrieve user identity information.
// @Tags Users
// @Produce json
// @Success 200 {object} IdentityResponse
// @Router /identity [get]
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
