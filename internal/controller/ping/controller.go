package ping

import (
    "github.com/labstack/echo/v4"
    "net/http"
)

func NewController(e *echo.Echo) {
    e.GET("/ping", func(c echo.Context) error {
        return c.JSON(http.StatusOK, "pong")
    })
}
