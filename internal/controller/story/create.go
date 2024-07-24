package story

import (
    "github.com/danieljuvito/iohub-server/internal/controller/middleware"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "github.com/labstack/echo/v4"
    "net/http"
)

type CreateRequest struct {
    Content string `json:"content"`
}

type CreateResponse struct {
    ID string `json:"id"`
}

// Create handles requests to create a new story.
//
// @Summary Create a story
// @Description Creates a new story with the provided content.
// @Tags Stories
// @Accept json
// @Produce json
// @Param request body CreateRequest true "Story content"
// @Success 201 {object} CreateResponse "Story created successfully"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Router /story [post]
func (c *controller) Create() {
    c.POST("", func(e echo.Context) error {
        ctx := e.Request().Context()

        session := e.Get("session").(model.Session)

        req := new(CreateRequest)
        if err := e.Bind(req); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        }

        res, err := c.storyService.Create(ctx, service.StoryCreateSpec{
            Story: &model.Story{
                UserID:  session.UserID,
                Content: req.Content,
            },
        })
        if err != nil {
            return err
        }
        return e.JSON(http.StatusOK, CreateResponse{
            ID: res.ID,
        })
    }, middleware.Auth)
}
