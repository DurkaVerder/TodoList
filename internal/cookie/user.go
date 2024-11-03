package cookie

import (
	"TodoList/internal/jwt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (c CookieManager) SaveJWTInCookie(token string, ctx echo.Context) error {

	cookie := new(http.Cookie)
	cookie.Name = "jwt-token"
	cookie.Value = token
	cookie.HttpOnly = true                          // Запрет доступа к cookie через JavaScript
	cookie.Expires = time.Now().Add(72 * time.Hour) // Устанавливаем время истечения cookie
	cookie.Path = "/"                               // Cookie будет доступно для всех путей сайта
	cookie.SameSite = http.SameSiteLaxMode          // Устанавливаем политику SameSite для предотвращения CSRF
	cookie.Secure = false                           // Используется только через HTTPS

	ctx.SetCookie(cookie)

	return nil
}

func (c CookieManager) GetJWTFromCookie(ctx echo.Context) (string, error) {
	Cookie, err := ctx.Cookie("jwt-token")
	if err != nil {
		return "", err
	}

	return Cookie.Value, nil
}

func (c CookieManager) GetUserIdByCookie(ctx echo.Context) (int, error) {
	token, err := c.GetJWTFromCookie(ctx)
	if err != nil {
		return -1, err
	}

	claims, err := jwt.ValidateJWT(token)
	if err != nil {
		return -1, err
	}
	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return -1, err
	}
	userId := int(userIdFloat)

	return userId, nil
}
