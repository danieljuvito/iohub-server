package story

import (
    "github.com/danieljuvito/iohub-server/internal/controller/middleware"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/labstack/echo/v4"
    "net/http"
)

type ListRequest struct {
}

type ListResponse struct {
    StoryID string `json:"story_id"`
}

func (c *controller) List() {
    c.GET("", func(e echo.Context) error {
        ctx := e.Request().Context()

        userID := e.QueryParam("user_id")

        res, err := c.storyService.List(ctx, service.StoryListSpec{
            UserID: userID,
        })
        if err != nil {
            return err
        }

        response := make([]ListResponse, 0)
        for _, data := range res.Data {
            response = append(response, ListResponse{
                StoryID: data.ID,
            })
        }

        return e.JSON(http.StatusOK, response)
    }, middleware.Auth)
}
