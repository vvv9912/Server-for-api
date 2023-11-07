package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"serverBot2/internal/model"
)

type Storager interface {
	SelectAllProducts(ctx context.Context) ([]model.Products, error)
}
type GetDB struct {
	Storage Storager
}

func NewGetDB(storage Storager) *GetDB {
	return &GetDB{Storage: storage}
}

func (s GetDB) GetDataDb(c echo.Context) error {

	products, err := s.Storage.SelectAllProducts(context.TODO())
	if err != nil {
		fmt.Println(err)
		return err
	}

	msg, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//f, _ := os.Create("output.json")
	//defer f.Close()
	//as_json, _ := json.MarshalIndent(products, "", "\t")
	//f.Write(as_json)
	//fmt.Println(products)
	byteArr := []byte(msg)
	return c.JSONBlob(http.StatusOK, byteArr)

	//return c.String(http.StatusBadRequest, "")
}
