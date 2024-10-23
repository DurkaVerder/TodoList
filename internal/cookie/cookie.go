package cookie

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

type Cookie interface {
	AddJWTInCookie(ctx echo.Context) error
}

type CookieManager struct {
	store *sessions.CookieStore
}

func NewCookie() Cookie {
	return CookieManager{sessions.NewCookieStore([]byte("key"))}
}


