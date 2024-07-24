package user

import (
    "github.com/danieljuvito/iohub-server/internal/controller/middleware"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "github.com/labstack/echo/v4"
    "net/http"
    "strconv"
)

type ListRequest struct {
}

type ListResponse struct {
    ID         string `json:"id"`
    Name       string `json:"name"`
    IsFollowed bool   `json:"is_followed"`
}

// List handles GET requests for retrieving a list of users.
// @Summary Get a list of users
// @Description Retrieve a list of users.
// @Tags Users
// @Produce json
// @Success 200 {array} ListResponse
// @Router /users [get]
func (c *controller) List() {
    c.GET("", func(e echo.Context) error {
        ctx := e.Request().Context()

        session := e.Get("session").(model.Session)

        search := e.QueryParam("search")
        page, _ := strconv.Atoi(e.QueryParam("page"))
        limit, _ := strconv.Atoi(e.QueryParam("limit"))

        res, err := c.userService.List(ctx, service.UserListSpec{
            UserID: session.UserID,
            Name:   search,
            Page:   page,
            Limit:  limit,
        })
        if err != nil {
            return err
        }

        response := make([]ListResponse, 0)
        for _, data := range res.Data {
            response = append(response, ListResponse{
                ID:         data.User.ID,
                Name:       data.Name,
                IsFollowed: data.Followee != nil,
            })
        }

        return e.JSON(http.StatusOK, response)
    }, middleware.Auth)
}
