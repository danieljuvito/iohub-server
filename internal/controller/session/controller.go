package session

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/labstack/echo/v4"
)

type ErrorResponse error

type controller struct {
    *echo.Group
    sessionService service.Session
}

func NewController(
    echo *echo.Echo,
    sessionService service.Session,
) {
    c := controller{
        Group:          echo.Group("/sessions"),
        sessionService: sessionService,
    }
    c.LogIn()
    c.LogOut()
}
