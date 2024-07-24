package ping

import (
    "github.com/labstack/echo/v4"
    "net/http"
)

type controller struct {
    *echo.Echo
}

func NewController(echo *echo.Echo) {
    c := controller{
        Echo: echo,
    }
    c.Ping()
}

// Ping handles GET requests for a simple health check.
// @Summary Ping endpoint
// @Description Responds with "pong" as a health check.
// @Tags Health
// @Produce json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func (c *controller) Ping() {
    c.GET("/ping", func(c echo.Context) error {
        return c.JSON(http.StatusOK, "pong")
    })
}
