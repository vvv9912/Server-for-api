package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"serverBot2/internal/model"
	"strings"
)

type PostStorager interface {
	AddProduct(ctx context.Context, product model.Products) error
	ChangeProductByArticle(ctx context.Context, product model.Products) error
}
type PostDB struct {
	Storage PostStorager
}
type ProductsGet struct {
	Article     int      `json:"article,omitempty"`
	Catalog     string   `json:"catalog,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	PhotoUrl    []string `json:"photo_url,omitempty"`
	Price       float64  `json:"price,omitempty"`
	Length      int      `json:"length"`
	Width       int      `json:"width"`
	Height      int      `json:"height"`
	Weight      int      `json:"weight"`
}

func (s *PostDB) PostSaveAddBD(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println(body)
		return err
	}
	bodyString := string(body)
	_ = bodyString
	//fmt.Println(bodyString)
	exampleRequest := new([]model.Products)
	//var a []model.Products
	var a []ProductsGet
	err = json.Unmarshal(body, &a)
	if err != nil {
		fmt.Println(err)
		return err
	}
	photosBytes := make([][]byte, 0)
	for _, product := range a {
		for _, base64Img := range product.PhotoUrl {
			prefix := "data:image/jpeg;base64,"
			encodedString := strings.TrimPrefix(base64Img, prefix)
			data, _ := base64.StdEncoding.DecodeString(encodedString)
			photosBytes = append(photosBytes, data)
			// Теперь `data` содержит байтовый массив изображения
		}
	}
	s.Storage.AddProduct(context.TODO(), model.Products{
		Article:     16,
		Catalog:     a[0].Catalog,
		Name:        a[0].Name,
		Description: a[0].Description,
		PhotoUrl:    photosBytes,
		Price:       a[0].Price,
		Length:      a[0].Length,
		Width:       a[0].Width,
		Height:      a[0].Height,
		Weight:      a[0].Weight,
	})
	if err := c.Bind(exampleRequest); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(exampleRequest)

	return c.JSONBlob(
		http.StatusOK,
		[]byte(""),
	)

}

func (s *PostDB) PostChangeBD(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println(body)
		return err
	}
	bodyString := string(body)
	_ = bodyString
	//fmt.Println(bodyString)

	var a []ProductsGet
	err = json.Unmarshal(body, &a)
	if err != nil {
		fmt.Println(err)
		return err
	}
	photosBytes := make([][]byte, 0)
	for i, product := range a {
		for _, base64Img := range product.PhotoUrl {
			prefix := "data:image/jpeg;base64,"
			encodedString := strings.TrimPrefix(base64Img, prefix)
			data, _ := base64.StdEncoding.DecodeString(encodedString)
			photosBytes = append(photosBytes, data)
			// Теперь `data` содержит байтовый массив изображения
		}

		err = s.Storage.ChangeProductByArticle(context.TODO(), model.Products{
			Article:     a[i].Article,
			Catalog:     a[i].Catalog,
			Name:        a[i].Name,
			Description: a[i].Description,
			PhotoUrl:    photosBytes,
			Price:       a[i].Price,
			Length:      a[i].Length,
			Width:       a[i].Width,
			Height:      a[i].Height,
			Weight:      a[i].Weight,
		})
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	//if err := c.Bind(exampleRequest); err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	//fmt.Println(exampleRequest)

	return c.JSONBlob(
		http.StatusOK,
		[]byte(""),
	)

}
