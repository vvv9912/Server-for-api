package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s GetDB) GetDownloadDB(c echo.Context) error {

	products, err := s.Storage.SelectAllProducts(context.TODO())
	if err != nil {
		fmt.Println(err)
		return err
	}
	//создаем файл

	//архивируем

	//отправляем
	msg, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//fmt.Println(products)
	byteArr := []byte(msg)
	return c.JSONBlob(http.StatusOK, byteArr)

	//return c.String(http.StatusBadRequest, "")
}
