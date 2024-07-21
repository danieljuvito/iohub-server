package controller

import (
    "github.com/danieljuvito/iohub-server/internal/controller/ping"
    "github.com/danieljuvito/iohub-server/internal/controller/session"
    "github.com/danieljuvito/iohub-server/internal/controller/user"
    "github.com/danieljuvito/iohub-server/internal/service"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func NewController(e *echo.Echo, s *service.Service) {
    e.Use(middleware.Logger())

    ping.NewController(e)
    user.NewController(e, s.UserService)
    session.NewController(e, s.SessionService)
}
