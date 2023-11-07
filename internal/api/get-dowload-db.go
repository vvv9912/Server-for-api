package api

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mholt/archiver"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"serverBot2/internal/exel"
)

func (s GetDB) GetDownloadDB(c echo.Context) error {

	products, err := s.Storage.SelectAllProducts(context.TODO())
	if err != nil {
		fmt.Println(err)
		c.Response().WriteHeader(http.StatusBadRequest)
		return err
	}
	pathBD := "testBD"
	_, err = os.Stat(pathBD)
	if !os.IsNotExist(err) {

		err = os.RemoveAll(pathBD)
		if err != nil {
			fmt.Println(err)
			c.Response().WriteHeader(http.StatusBadRequest)
			return err
		}
	}
	err = os.Mkdir(pathBD, os.ModePerm)
	ex2 := exel.NewExcel(path.Join(pathBD, "test.xlsx"))
	//
	f := excelize.NewFile()
	ex2.CreateFile(f)
	ex2.WriteStamp("Catalog")
	for i := range products {
		var photos []string
		pathCatalog := path.Join(pathBD, products[i].Catalog)
		if len(products[i].PhotoUrl) != 0 {
			err = os.Mkdir(pathCatalog, os.ModePerm)
			if err != nil {
				fmt.Println(err)
				c.Response().WriteHeader(http.StatusBadRequest)
				return err
			}
		}

		for k := range products[i].PhotoUrl {
			fileName := products[i].Catalog + "_" + products[i].Name + "_" + generateRandomName()
			err := ioutil.WriteFile(path.Join(pathCatalog, fileName), products[i].PhotoUrl[k], 0644)
			if err != nil {
				fmt.Println(err)
				c.Response().WriteHeader(http.StatusBadRequest)
				return err
			}
			photos = append(photos, "/"+products[i].Catalog+"/"+fileName)
		}

		ex2.Write(i+2, exel.ProductsPars{
			Article:     products[i].Article,
			Catalog:     products[i].Catalog,
			Name:        products[i].Name,
			Description: products[i].Description,
			PhotoUrl:    photos,
			Price:       products[i].Price,
			Length:      products[i].Length,
			Width:       products[i].Width,
			Height:      products[i].Height,
			Weight:      products[i].Weight,
		}, "Catalog")
	}

	ex2.CloseFile()

	fileZip := "bd.zip"
	_, err = os.Stat(fileZip)
	if !os.IsNotExist(err) {

		err = os.RemoveAll(fileZip)
		if err != nil {
			fmt.Println(err)
			c.Response().WriteHeader(http.StatusBadRequest)
			return err
		}
	}
	err = archiver.Archive([]string{pathBD}, fileZip)

	if err != nil {
		fmt.Println(err)
		c.Response().WriteHeader(http.StatusBadRequest)
		return err
	}
	_, err = os.Stat(pathBD)
	if !os.IsNotExist(err) {

		err = os.RemoveAll(pathBD)
		if err != nil {
			fmt.Println(err)
			c.Response().WriteHeader(http.StatusBadRequest)
			return err
		}
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/zip")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+fileZip)
	c.Response().WriteHeader(http.StatusOK)

	return c.File(fileZip)

}
func generateRandomName() string {
	// Генерируем случайное имя для файла
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(randomBytes) + ".jpg"
}
