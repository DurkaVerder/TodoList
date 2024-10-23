package cookie

import "github.com/labstack/echo"

func (c CookieManager) AddJWTInCookie(ctx echo.Context) error {
	session, err := c.store.Get(ctx.Request(), "cookie-todolist")
	if err != nil {
		return err
	}

	// Other logical

	if err := session.Save(ctx.Request(), ctx.Response()); err != nil {
		return err
	}

	return nil
}
