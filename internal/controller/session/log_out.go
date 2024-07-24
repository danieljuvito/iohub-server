package session

import (
    "github.com/danieljuvito/iohub-server/internal/controller/middleware"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "github.com/labstack/echo/v4"
    "net/http"
)

// LogOut handles user logout requests.
//
// @Summary User logout
// @Description Logs out the currently authenticated user.
// @Tags Sessions
// @Security ApiKeyAuth
// @Success 204 "No content"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Router /log-out [delete]
func (c *controller) LogOut() {
    c.DELETE("/log-out", func(e echo.Context) error {
        ctx := e.Request().Context()

        session, ok := e.Get("session").(model.Session)
        if !ok {
            return echo.NewHTTPError(http.StatusUnauthorized, "session not found")
        }

        _, err := c.sessionService.LogOut(ctx, service.SessionLogOutSpec{
            ID: session.ID,
        })
        if err != nil {
            return err
        }

        return e.JSON(http.StatusOK, "OK")
    }, middleware.Auth)
}
