package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	// "github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func calc_timeLimitSecs() int {
	now_time := time.Now()
	y := now_time.Year()
	m := now_time.Month()
	d := now_time.AddDate(0, 0, 1).Day()
	all_day := time.Date(y, m, d, 0, 0, 0, 0, now_time.Location())

	diff_secs := all_day.Unix() - now_time.Unix() // UnixTimeによる差分取得(単位: 秒)
	return int(diff_secs)                         // int64だとMaxAgeに渡せない
}

func setCookie(c echo.Context) error {
	post := new(ReceiveLoginInfo)
	err := c.Bind(post)
	if err != nil {
		return c.String(http.StatusCreated, "Missed!")
	}

	msg := []byte(post.Student_id)
	cookie := new(http.Cookie)
	cookie.Name = "student_id"
	cookie.Value = base64.URLEncoding.EncodeToString(msg)
	cookie.Expires = time.Now().Add(10 * time.Second) // 10secs
	cookie.MaxAge = 10 // 10secs
	cookie.Path = "/"
	cookie.Secure = false
	// cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode

	c.SetCookie(cookie)
	fmt.Println("~~set cookie~~")
	fmt.Println(cookie)
	return c.String(http.StatusCreated, "Create new Cookie!")
}

func readCookie(c echo.Context) error {
	/* post := ReceiveLoginInfo{}
	err := c.Bind(post)
	if err != nil {
		return c.String(http.StatusCreated, "Bind Missed!")
	} */

	cookie, err := c.Cookie("student_id")
	fmt.Println("~~read~~")
	fmt.Println(cookie)
	//fmt.Println(cookie.Name)
	//fmt.Println(cookie.Value)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusCreated, "Cookie Missed!")
	}
	return c.String(http.StatusCreated, "Success!")
}

func main() {
	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	// ミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins: []string{"http://localhost:3000"},
		/* AllowHeaders: []string{
			echo.HeaderAccept,
			echo.HeaderContentType,
			echo.HeaderAccessControlAllowCredentials,
			echo.HeaderOrigin,
		}, */
		// 今回はPOSTメソッドのみ使用する
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
	}))

	// セッションによる自動ログイン
	e.GET("/readCookie", readCookie)
	e.POST("/setCookie", setCookie)

	// APIを叩いてデータベースと連携
	e.POST("/attendance/attend", PostAttend) // 出欠情報を送る
	e.POST("/attendance/login", PostLogin)   // 本人確認の実施
	e.POST("/attendance/reload", PostReload)

	e.Logger.Fatal(e.Start(":8080"))
}
