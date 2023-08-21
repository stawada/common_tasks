package main

import (
	"fmt"
	"net/http"
	"time"

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

	cookie := &http.Cookie{
		Name:   "student_id",
		Value:  post.Student_id + post.Hashed_password,
		MaxAge: calc_timeLimitSecs(), // 翌日0:00まで
		Path:   "/",
	}

	c.SetCookie(cookie)
	return c.String(http.StatusCreated, "Create new Cookie!")
}

func readCookie(c echo.Context) error {
	cookie, err1 := c.Cookie("student_id")

	resJson := ReturnLoginInfo{}
	search_studentId := fmt.Sprintf("SELECT student_id from student WHERE student_id='%s';", cookie.Value)
	var student_id string
	err2 := db.QueryRow((search_studentId)).Scan(&student_id)

	if err1 != nil || err2 != nil {
		resJson.Match_flag = 0
	} else {
		resJson.Match_flag = 1
	}
	resJson.Http_status = http.StatusCreated

	return c.JSON(http.StatusCreated, resJson)
}

func main() {
	e := echo.New()

	// ミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:3000"},
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
