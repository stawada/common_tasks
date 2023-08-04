package main

import (
	"fmt"
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

/* 出欠情報用のjson */
type Attend struct {
	Student_id   string `json:"student_id"`
	Student_name string `json:"student_name"`
	Hashed_pw    string `json:"hashed_pw"`
}

/* 返り値用のjson */
type ReturnJson struct {
	Attend_flag int `json:"attend_flag"`
	Http_status int `json:"http_status"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "host=localhost user=stawada dbname=sample sslmode=disable")
	if err != nil {
		panic("Not found Database.")
	}
}

func PostFromJson(c echo.Context) error {
	post := new(Attend)

	err := c.Bind(post)
	if err != nil {
		return err
	}

	res := Attend{}
	resJson := ReturnJson{}

	select_sentence := fmt.Sprintf("SELECT student_id, hashed_pw FROM student where student_id='%s' and hashed_pw='%s'", post.Student_id, post.Hashed_pw)
	if err := db.QueryRow(select_sentence).Scan(&res.Student_id, &res.Hashed_pw); err != nil {
		// 失敗時はフラグ=0
		resJson.Attend_flag = 0
		resJson.Http_status = http.StatusCreated
	} else {
		// 成功時はフラグ=1
		resJson.Attend_flag = 1
		resJson.Http_status = http.StatusCreated
	}

	return c.JSON(http.StatusCreated, resJson)
}
