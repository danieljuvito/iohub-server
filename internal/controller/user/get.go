package user

import (
    "github.com/danieljuvito/iohub-server/internal/controller/middleware"
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/labstack/echo/v4"
    "net/http"
)

type GetResponse struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

// Get handles GET requests for retrieving user details.
// @Summary Get user details
// @Description Retrieve user details by ID.
// @Tags Users
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} GetResponse
// @Router /users/{id} [get]
func (c *controller) Get() {
    c.GET("/:id", func(e echo.Context) error {
        ctx := e.Request().Context()

        id := e.Param("id")

        res, err := c.userService.Get(ctx, service.UserGetSpec{
            ID: id,
        })
        if err != nil {
            return err
        }
        return e.JSON(http.StatusOK, GetResponse{
            ID:   res.ID,
            Name: res.Name,
        })
    }, middleware.Auth)
}
