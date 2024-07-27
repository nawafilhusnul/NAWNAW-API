package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	auth "github.com/nawafilhusnul/NAWNAW-API/auth/usecase"
	"github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
	"github.com/nawafilhusnul/NAWNAW-API/model"
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
		req := &model.LoginRequest{}

		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, response.NewResponse().WithError(err))
		}

		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, response.NewResponse().WithError(err))
		}

		user, err := h.uc.Login(ctx, req.Identifier, req.Password, req.Timezone)
		if err != nil {
			return c.JSON(response.GetErrorStatusCode(err), response.NewResponse().WithError(err))
		}

		return c.JSON(http.StatusOK, response.NewResponse().WithData(user, "Login success"))
	})
}

func (h *handler) Register() echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		ctx := c.(*ctx.Ctx)
		req := &model.Auth{}

		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, response.NewResponse().WithError(err))
		}

		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, response.NewResponse().WithError(err))
		}

		err := h.uc.Register(ctx, req)
		if err != nil {
			return c.JSON(response.GetErrorStatusCode(err), response.NewResponse().WithError(err))
		}

		return c.JSON(http.StatusOK, response.NewResponse().WithData(map[string]interface{}{
			"id": req.ID,
		}, "Register success"))
	})
}

func (h *handler) GetOne() echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		ctx := c.(*ctx.Ctx)

		id := ctx.GetUser().UserID
		user, err := h.uc.GetOne(ctx, id)
		if err != nil {
			return c.JSON(response.GetErrorStatusCode(err), response.NewResponse().WithError(err))
		}

		return c.JSON(http.StatusOK, response.NewResponse().WithData(user, "Get one success"))
	})
}
