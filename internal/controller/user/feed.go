package user

import (
    "github.com/danieljuvito/iohub-server/internal/controller/middleware"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "github.com/labstack/echo/v4"
    "net/http"
    "strconv"
)

type FeedRequest struct {
}

type FeedResponse struct {
    UserID     string `json:"user_id"`
    Name       string `json:"name"`
    StoryCount int    `json:"story_count"`
    IsFollowed bool   `json:"is_followed"`
}

func (c *controller) Feed() {
    c.GET("/feed", func(e echo.Context) error {
        ctx := e.Request().Context()

        session := e.Get("session").(model.Session)

        page, _ := strconv.Atoi(e.QueryParam("page"))
        limit, _ := strconv.Atoi(e.QueryParam("limit"))

        res, err := c.userService.Feed(ctx, service.UserFeedSpec{
            UserID: session.UserID,
            Page:   page,
            Limit:  limit,
        })
        if err != nil {
            return err
        }

        response := make([]FeedResponse, 0)
        for _, data := range res.Data {
            response = append(response, FeedResponse{
                UserID:     data.User.ID,
                Name:       data.Name,
                StoryCount: len(data.Stories),
                IsFollowed: true,
            })
        }

        return e.JSON(http.StatusOK, response)
    }, middleware.Auth)
}
