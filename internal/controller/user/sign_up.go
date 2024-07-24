package user

import (
    "github.com/danieljuvito/iohub-server/internal/domain/interface/service"
    "github.com/danieljuvito/iohub-server/internal/domain/model"
    "github.com/labstack/echo/v4"
    "net/http"
)

type SignUpRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type SignUpResponse struct {
    ID string `json:"id"`
}

// SignUp handles POST requests for user sign-up.
// @Summary User sign-up
// @Description Register a new user.
// @Tags Users
// @Accept json
// @Produce json
// @Param request body SignUpRequest true "Sign-up request"
// @Success 201 {object} SignUpResponse
// @Router /sign-up [post]
func (c *controller) SignUp() {
    c.POST("/sign-up", func(e echo.Context) error {
        ctx := e.Request().Context()

        req := new(SignUpRequest)
        if err := e.Bind(req); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        }

        res, err := c.userService.SignUp(ctx, service.UserSignUpSpec{
            User: &model.User{
                Name:     req.Name,
                Email:    req.Email,
                Password: req.Password,
            },
        })
        if err != nil {
            return err
        }
        return e.JSON(http.StatusOK, SignUpResponse{
            ID: res.ID,
        })
    })
}
