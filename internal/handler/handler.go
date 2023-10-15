package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthHandler(c echo.Context) error {

	err := c.Render(http.StatusOK, "auth.html", map[string]interface{}{
		"name":    "AUTH",
		"CSSPath": "/static/css/styleauth.css",
	})
	fmt.Println(err)
	return err
}
func AuthHandler2(c echo.Context) error {

	err := c.Render(http.StatusOK, "auth2.html", map[string]interface{}{
		"name":    "AUTH",
		"CSSPath": "/static/css/styleauth2.css",
	})
	fmt.Println(err)
	return err
}
func AuthHandler3(c echo.Context) error {

	err := c.Render(http.StatusOK, "auth3.html", map[string]interface{}{
		"name":    "AUTH",
		"CSSPath": "/static/css/styleauth3.css",
	})
	fmt.Println(err)
	return err
}
func BdHandler(c echo.Context) error {

	err := c.Render(http.StatusOK, "bd.html", map[string]interface{}{
		"name":    "AUTH",
		"CSSPath": "/static/css/style_bd.css",
	})
	fmt.Println(err)
	return err
}
