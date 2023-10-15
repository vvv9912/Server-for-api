package api

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"serverBot2/internal/model"
)

func sha256Hash(input string) string {
	// Создаем новый хеш SHA-256
	hasher := sha256.New()

	// Преобразуем строку в байты и передаем хеш-функции
	hasher.Write([]byte(input))

	// Получаем хеш в виде среза байтов
	hashBytes := hasher.Sum(nil)

	// Преобразуем срез байтов в строку в шестнадцатеричном формате
	hashedString := hex.EncodeToString(hashBytes)

	return hashedString
}
func PostAuth(c echo.Context) error {
	//body, err := ioutil.ReadAll(c.Request().Body)
	//if err != nil {
	//	fmt.Println(body)
	//	return err
	//}
	//bodyString := string(body)
	//fmt.Println(bodyString)
	exampleRequest := new(model.ExampleRequest)
	if err := c.Bind(exampleRequest); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(exampleRequest)
	cookie := new(http.Cookie)

	cookie.Name = "login"
	cookie.Value = "kakoitotokenalal"
	c.SetCookie(cookie)
	//c.String(http.StatusOK, "set cookie")
	var resp bytes.Buffer
	var b = []byte(
		fmt.Sprintf(`{
      "success": %s
    }`, "true"),
	)

	err := json.Indent(&resp, b, "", "  ")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSONBlob(
		http.StatusOK,
		[]byte(
			fmt.Sprintf("%s", resp.Bytes()),
		),
	)
	//return c.JSONBlob(
	//	http.StatusOK,
	//	[]byte(
	//		fmt.Sprintf("%s", resp.Bytes()),
	//	),
	//)
	//return c.String(http.StatusBadRequest, "")
}
