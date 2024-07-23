package controller

import (
    "github.com/danieljuvito/iohub-server/internal/controller/ping"
    "github.com/danieljuvito/iohub-server/internal/controller/session"
    "github.com/danieljuvito/iohub-server/internal/controller/user"
    "github.com/danieljuvito/iohub-server/internal/service"
    "github.com/danieljuvito/iohub-server/internal/util/errorutil"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "net/http"
)

func NewController(e *echo.Echo, s *service.Service) {
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
        AllowMethods: []string{"*"},
    }))
    e.Use(middleware.Recover())
    e.Use(middleware.Logger())

    e.HTTPErrorHandler = NewHttpErrorHandler(map[error]int{
        errorutil.Invalid:      http.StatusBadRequest,
        errorutil.AlreadyExist: http.StatusConflict,
        errorutil.NotFound:     http.StatusNotFound,
        errorutil.Unauthorized: http.StatusUnauthorized,
    }).Handler

    ping.NewController(e)
    user.NewController(e, s.UserService)
    session.NewController(e, s.SessionService)
}
