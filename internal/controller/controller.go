package controller

import (
    "fmt"
    "github.com/danieljuvito/iohub-server/internal/controller/ping"
    "github.com/danieljuvito/iohub-server/internal/controller/session"
    "github.com/danieljuvito/iohub-server/internal/controller/story"
    "github.com/danieljuvito/iohub-server/internal/controller/user"
    "github.com/danieljuvito/iohub-server/internal/service"
    "github.com/danieljuvito/iohub-server/internal/util/errorutil"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    echoSwagger "github.com/swaggo/echo-swagger"
    "golang.org/x/net/websocket"
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

    e.GET("/ws", hello)
    e.GET("/swagger/*", echoSwagger.WrapHandler)
    ping.NewController(e)
    user.NewController(e, s.UserService)
    session.NewController(e, s.SessionService)
    story.NewController(e, s.StoryService)
}

func hello(c echo.Context) error {
    websocket.Handler(func(ws *websocket.Conn) {
        defer ws.Close()
        for {
            // Write
            err := websocket.Message.Send(ws)
            if err != nil {
                c.Logger().Error(err)
            }

            // Read
            msg := ""
            err = websocket.Message.Receive(ws, &msg)
            if err != nil {
                c.Logger().Error(err)
            }
            fmt.Printf("%s\n", msg)
        }
    }).ServeHTTP(c.Response(), c.Request())
    return nil
}
