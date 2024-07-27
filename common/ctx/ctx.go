package ctx

import (
	"context"

	"github.com/labstack/echo/v4"
)

type ContextUser struct {
	UserID    int
	Roles     map[string]bool
	Platforms map[string]bool
	Timezone  string
}

type Ctx struct {
	echo.Context
	User *ContextUser
}

// SetUser sets the user information in the custom context.
// Usage example:
//
//	user := &ContextUser{UserID: 1, Roles: map[string]bool{"admin": true}, Platforms: map[string]bool{"web": true}, Timezone: "UTC"}
//	ctx.SetUser(user)
func (c *Ctx) SetUser(user *ContextUser) {
	c.User = user
}

// GetUser retrieves the user information from the custom context.
// Usage example:
//
//	user := ctx.GetUser()
//	fmt.Println(user.UserID)
func (c *Ctx) GetUser() *ContextUser {
	return c.User
}

// RequestContext returns the underlying request context.
// Usage example:
//
//	reqCtx := ctx.RequestContext()
//	fmt.Println(reqCtx)
func (c *Ctx) RequestContext() context.Context {
	return c.Context.Request().Context()
}

// NewCtx creates a new custom context and passes it to the next handler in the chain.
// Usage example:
//
//	e := echo.New()
//	e.Use(NewCtx)
//	e.GET("/", func(c echo.Context) error {
//	    ctx := c.(*Ctx)
//	    return ctx.String(http.StatusOK, "Hello, World!")
//	})
func NewCtx(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &Ctx{Context: c}
		return next(cc)
	}
}
