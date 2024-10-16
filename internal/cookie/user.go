package cookie

import "github.com/labstack/echo"

func (c CookieManager) AddUserIdInCookie(ctx echo.Context, userId int) error {
	session, err := c.store.Get(ctx.Request(), "cookie-todolist")
	if err != nil {
		return err
	}

	session.Values["userId"] = userId

	if err := session.Save(ctx.Request(), ctx.Response()); err != nil {
		return err
	}

	return nil
}
