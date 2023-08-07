/* ファイルサーバーの構築 */

package main

import (
	"net/http"

	"github.com/labstack/echo/v4" // インストール必要
)

func customHTTPErrorHandler(err error, c echo.Context) {
	/* HTTPの接続チェック */

	code := http.StatusInternalServerError // サーバー接続の結果(フラグ？)を取得
	msg := http.StatusText(code)           // テキストで結果を保持

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	}

	if code == http.StatusNotFound {
		c.File("./404.html")
		return
	}

	c.JSON(code, map[string]interface{}{
		"error": map[string]interface{}{
			"code":    code,
			"message": msg,
		},
	})
}

func getGreet(c echo.Context) error {
	/* 対象のファイルを出力する */

	// 相対パスだとうまくいかないので絶対パスに変換
	target_path := "./index.html"
	return c.File(target_path)
}

func main() {
	e := echo.New() // インスタンス生成

	e.HTTPErrorHandler = customHTTPErrorHandler // 404用のページを出力するやつを渡す

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Pasona!")
	})
	e.GET("/pasona", getGreet)
	e.Logger.Fatal(e.Start(":8080"))
}
