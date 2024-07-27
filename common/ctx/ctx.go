package ctx

import (
	"context"

	"github.com/labstack/echo/v4"
)

type ContextUser struct {
	Name string
}

type Ctx struct {
	echo.Context
	User *ContextUser
}

func (c *Ctx) SetUser(user *ContextUser) {
	c.User = user
}

func (c *Ctx) GetUser() *ContextUser {
	return c.User
}

func (c *Ctx) RequestContext() context.Context {
	return c.Context.Request().Context()
}

func NewCtx(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &Ctx{Context: c}
		return next(cc)
	}
}
