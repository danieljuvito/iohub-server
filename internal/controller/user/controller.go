package user

import (
    "github.com/danieljuvito/iohub-server/internal/interface/service"
    "github.com/labstack/echo/v4"
)

type controller struct {
    *echo.Group
    userService service.User
}

func NewController(
    echo *echo.Echo,
    userService service.User,
) {
    c := controller{
        Group:       echo.Group("/users"),
        userService: userService,
    }
    c.SignUp()
}
