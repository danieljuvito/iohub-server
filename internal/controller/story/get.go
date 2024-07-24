package story

import (
    "github.com/danieljuvito/iohub-server/internal/controller/middleware"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/labstack/echo/v4"
    "net/http"
)

type GetResponse struct {
    ID      string `json:"id"`
    Content string `json:"content"`
}

func (c *controller) Get() {
    c.GET("/:id", func(e echo.Context) error {
        ctx := e.Request().Context()

        id := e.Param("id")

        res, err := c.storyService.Get(ctx, service.StoryGetSpec{
            ID: id,
        })
        if err != nil {
            return err
        }
        return e.JSON(http.StatusOK, GetResponse{
            ID:      res.ID,
            Content: res.Content,
        })
    }, middleware.Auth)
}
