package api

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Data struct {
	Id       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func GetData(c echo.Context) error {
	a := []Data{
		{
			Id:       "1",
			Email:    "1",
			Password: "1",
		},
		{
			Id:       "2",
			Email:    "2",
			Password: "2",
		},
		{
			Id:       "3",
			Email:    "3",
			Password: "3",
		},
	}
	msg, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return err
	}
	byteArr := []byte(msg)
	return c.JSONBlob(http.StatusOK, byteArr)

	//return c.String(http.StatusBadRequest, "")
}
