package cookie

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

type Cookie interface {
	AddUserIdInCookie(ctx echo.Context, userId int) error
}

type CookieManager struct {
	store *sessions.CookieStore
}

func NewCookie() Cookie {
	return CookieManager{sessions.NewCookieStore([]byte("key"))}
}


