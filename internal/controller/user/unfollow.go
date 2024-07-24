package user

import (
    "github.com/danieljuvito/iohub-server/internal/controller/middleware"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "github.com/labstack/echo/v4"
    "net/http"
)

// Unfollow handles PATCH requests to unfollow a user.
// @Summary Unfollow a user
// @Description Unfollow a specific user by ID.
// @Tags Users
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {string} string "Successfully unfollowed user"
// @Router /unfollow/{id} [patch]
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
