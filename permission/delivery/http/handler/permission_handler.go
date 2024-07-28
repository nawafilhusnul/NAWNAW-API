package permission

import (
	"net/http"

	"github.com/labstack/echo/v4"
	cc "github.com/nawafilhusnul/NAWNAW-API/common/ctx"
	"github.com/nawafilhusnul/NAWNAW-API/common/datatypes"
	"github.com/nawafilhusnul/NAWNAW-API/common/response"
	"github.com/nawafilhusnul/NAWNAW-API/model"
	permission "github.com/nawafilhusnul/NAWNAW-API/permission/usecase"
)

type handler struct {
	uc permission.Usecase
}

func New(uc permission.Usecase) *handler {
	return &handler{uc: uc}
}

func (h *handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*cc.Ctx)
		var req model.CreatePermissionRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.NewResponse().WithError(err))
		}

		permission := &model.Permission{
			Name:     datatypes.SetNullString(req.Name),
			ModuleID: datatypes.ID(req.ModuleID),
		}
		err := h.uc.Create(ctx, permission)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.NewResponse().WithError(err))
		}

		return c.JSON(http.StatusCreated, response.NewResponse().WithData(permission.ID, "Permission created successfully"))
	}
}
