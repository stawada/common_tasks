package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		return
	}

	c.JSON(code, map[string]interface{}{
		"error": map[string]interface{}{
			"code":    code,
			"message": msg,
		},
	})
}

type ReceiveReload struct {
	Student_id string `json:"student_id"`
	Now_time   int    `json:"now_time"`
}
type ReturnReload struct {
	Subject_name string `json:"subject_name"`
	Http_status  int    `json:"http_status"`
}

func main() {
	e := echo.New() // インスタンス生成

	// e.File("/", "../../vue-test-project/dist")
	e.Static("/", "../../vue-test-project/dist")

	e.Logger.Fatal(e.Start(":3000"))
	e.HTTPErrorHandler = customHTTPErrorHandler // 404用のページを出力するやつを渡す

	url := "http://localhost:8080/attendance/reload"
	client := http.Client{}
	// ヘッダー情報
	header := http.Header{}
	header.Add("Content-Type", "application/json")
	header.Add("Accept", "application/json")
	// ボディ情報
	body := ReceiveReload{
		Student_id: "1000000000",
		Now_time:   1691130450,
	}
	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header = header
	resp, _ := client.Do(req)

	// 読み込み
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bytes := []byte(byteArray)
	var response ReturnReload
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		panic(err)
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(response)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	// 閉じる
	resp.Body.Close()

	// e.Logger.Fatal(e.Start(":3000"))
}
