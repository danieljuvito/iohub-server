package user

import (
    "github.com/danieljuvito/iohub-server/internal/controller/middleware"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "github.com/labstack/echo/v4"
    "net/http"
)

// Follow handles PATCH requests to follow a user.
// @Summary Follow a user
// @Description Follow a specific user by ID.
// @Tags Users
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {string} string "Successfully followed user"
// @Router /follow/{id} [patch]
func (c *controller) Follow() {
    c.PATCH("/follow/:id", func(e echo.Context) error {
        ctx := e.Request().Context()

        id := e.Param("id")

        session := e.Get("session").(model.Session)

        _, err := c.userService.Follow(ctx, service.UserFollowSpec{
            UserID:         session.UserID,
            FolloweeUserID: id,
        })
        if err != nil {
            return err
        }
        return e.JSON(http.StatusOK, "OK")
    }, middleware.Auth)
}
