package mw

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

// токен(key) -> id
type Cacher interface {
	GetToken(token string) (string, bool)
	NewToken(Token string, ID string) error
}
type MW struct {
	c Cacher
}

func NewMW(c Cacher) *MW {
	return &MW{c: c}
}

func (m *MW) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//if len(c.Cookies()) == 0 {
		//	err := next(c)
		//}
		if len(c.Cookies()) != 0 {
			token, err := c.Cookie("token")
			_ = token
			if err != nil {
				fmt.Println(err)
				return next(c)
			} else {

			}

		}
		return nil
	}
}
