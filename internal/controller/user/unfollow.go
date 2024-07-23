package user

import (
    "github.com/danieljuvito/iohub-server/internal/controller/middleware"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "github.com/labstack/echo/v4"
    "net/http"
)

func (c *controller) Unfollow() {
    c.PATCH("/unfollow/:id", func(e echo.Context) error {
        ctx := e.Request().Context()

        id := e.Param("id")

        session := e.Get("session").(model.Session)

        _, err := c.userService.Unfollow(ctx, service.UserUnfollowSpec{
            UserID:         session.UserID,
            FolloweeUserID: id,
        })
        if err != nil {
            return err
        }
        return e.JSON(http.StatusOK, "OK")
    }, middleware.Auth)
}
