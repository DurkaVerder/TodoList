package cookie

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

type Cookie interface {
	SaveJWTInCookie(token string, ctx echo.Context) error
	GetJWTFromCookie(ctx echo.Context) (string, error)
}

type CookieManager struct {
	store *sessions.CookieStore
}

func NewCookie() Cookie {
	return CookieManager{sessions.NewCookieStore([]byte("key"))}
}
