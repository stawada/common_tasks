package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// ミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		// 今回はPOSTメソッドのみ使用する
		AllowMethods: []string{
			// http.MethodGet,
			// http.MethodPut,
			http.MethodPost,
			// http.MethodDelete,
		},
	}))

	// APIを叩いてデータベースと連携
	e.POST("/attendance/attend", PostAttend) // 出欠情報を送る
	e.POST("/attendance/login", PostLogin)   // 本人確認の実施
	e.POST("/attendance/reload", PostReload)

	e.Logger.Fatal(e.Start(":8080"))
}
