package user

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/labstack/echo/v4"
    "net/http"
)

type ListRequest struct {
}

type ListResponse struct {
    ID    string `json:"id"`
    Email string `json:"email"`
}

func (c *controller) List() {
    c.GET("/", func(e echo.Context) error {
        ctx := e.Request().Context()

        search := e.Param("search")

        res, err := c.userService.List(ctx, service.UserListSpec{
            Search: search,
        })
        if err != nil {
            return err
        }

        var response []ListResponse
        for _, data := range res.Data {
            response = append(response, ListResponse{
                ID:    data.ID,
                Email: data.Email,
            })
        }

        return e.JSON(http.StatusOK, response)
    })
}
