package module

import (
	"net/http"

	"github.com/labstack/echo/v4"
	cc "github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
	module "github.com/nawafilhusnul/NAWNAW-API/module/usecase"
)

type handler struct {
	uc module.Usecase
}

func New(uc module.Usecase) *handler {
	return &handler{uc: uc}
}

func (h *handler) FindAll() echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		ctx := c.(*cc.Ctx)

		modules, err := h.uc.FindAll(ctx)
		if err != nil {
			return c.JSON(response.GetErrorStatusCode(err), response.NewResponse().WithError(err))
		}

		return c.JSON(http.StatusOK, response.NewResponse().WithData(modules, "Modules fetched successfully"))
	})
}
