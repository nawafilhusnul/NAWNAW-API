package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	auth "github.com/nawafilhusnul/big-app/auth/usecase"
	"github.com/nawafilhusnul/big-app/common/ctx"
	"github.com/nawafilhusnul/big-app/common/response"
	"github.com/nawafilhusnul/big-app/model"
)

type handler struct {
	uc auth.Usecase
}

func NewAuthHandler(uc auth.Usecase) *handler {
	return &handler{uc: uc}
}

func (h *handler) Login() echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		ctx := c.(*ctx.Ctx)
		h.uc.Login(ctx, "email", "password")
		return nil
	})
}

func (h *handler) Register() echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		ctx := c.(*ctx.Ctx)
		req := &model.Auth{}

		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		err := h.uc.Register(ctx, req)
		if err != nil {
			return c.JSON(http.StatusOK, response.NewResponse().WithData(1, "success"))
			// return c.JSON(response.GetErrorStatusCode(err), response.NewResponse().WithError(err))
		}

		return c.JSON(http.StatusOK, response.NewResponse().WithData("OK!", "Register success"))
	})
}
