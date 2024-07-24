package story

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/labstack/echo/v4"
)

type controller struct {
    *echo.Group
    storyService service.Story
}

func NewController(
    echo *echo.Echo,
    storyService service.Story,
) {
    c := controller{
        Group:        echo.Group("/stories"),
        storyService: storyService,
    }
    c.Get()
    c.List()
    c.Create()
}
