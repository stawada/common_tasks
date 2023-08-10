package main

import (
	"database/sql"
	"fmt"

	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "host=localhost user=stawada dbname=sample sslmode=disable")
	if err != nil {
		panic("Not found Database.")
	}
}

/* ログイン用API */
type ReceiveLoginInfo struct {
	Student_id      string `json:"student_id"`
	Hashed_password string `json:"hashed_password"`
}

type ReturnLoginInfo struct {
	Match_flag  int `json:"match_flag"`
	Http_status int `json:"http_status"`
}

func PostLogin(c echo.Context) error {
	post := new(ReceiveLoginInfo)
	err := c.Bind(post)
	if err != nil {
		return err
	}

	checkJson := ReceiveLoginInfo{}
	resJson := ReturnLoginInfo{}
	select_sentence := fmt.Sprintf("SELECT student_id, hashed_password FROM student where student_id='%s' and hashed_password='%s';", post.Student_id, post.Hashed_password)
	fmt.Println(select_sentence)
	if err := db.QueryRow(select_sentence).Scan(&checkJson.Student_id, &checkJson.Hashed_password); err != nil && &checkJson != nil {
		// 失敗時はフラグ=0
		resJson.Match_flag = 0
		resJson.Http_status = http.StatusCreated
	} else {
		// 成功時はフラグ=1
		resJson.Match_flag = 1
		resJson.Http_status = http.StatusCreated
	}

	return c.JSON(http.StatusCreated, resJson)
}

/* 欠席情報用API */
type ReceiveAttendInfo struct {
	Attend_flag int    `json:"attend_flag"`
	Student_id  string `json:"student_id"`
	Subject_id  string `json:"subject_id"`
}

type ReturnAttendInfo struct {
	Http_status int `json:"http_status"`
}

func PostAttend(c echo.Context) error {
	post := new(ReceiveAttendInfo)
	err := c.Bind(post)
	if err != nil {
		return err
	}

	resJson := ReturnAttendInfo{}
	update_sentence := fmt.Sprintf("UPDATE attendance_information SET attendance_flag=%d WHERE student_id='%s' AND lecture_history_id='%s';", post.Attend_flag, post.Student_id, post.Subject_id)
	if _, err := db.Exec(update_sentence); err != nil {
		resJson.Http_status = http.StatusCreated
	} else {
		resJson.Http_status = http.StatusCreated
	}

	return c.JSON(http.StatusCreated, resJson)
}

/* 講義情報送信用API */
type ReceiveReload struct {
	Student_id string `json:"student_id"`
	Now_time   int    `json:"now_time"`
}

type ReturnReload struct {
	Subject_name string `json:"subject_name"`
	Http_status  int    `json:"http_status"`
}

func PostReload(c echo.Context) error {
	post := new(ReceiveReload)
	err := c.Bind(post)
	if err != nil {
		return err
	}

	var where_phase string
	if post.Now_time == 0 {
		where_phase = fmt.Sprintf("WHERE attendance_information.student_id='%s';", post.Student_id)
	} else {
		where_phase = fmt.Sprintf("WHERE lecture_history.lecture_date_and_time-600<=%d AND lecture_history.lecture_date_and_time+600>=%d AND attendance_information.student_id='%s';", post.Now_time, post.Now_time, post.Student_id)
	}
	extract_sentence := `SELECT lecture_catalog.lecture_name FROM lecture_catalog 
						INNER JOIN lecture_history ON lecture_catalog.lecture_id=lecture_history.lecture_catalog_id 
						INNER JOIN attendance_information ON lecture_history.lecture_history_id=attendance_information.lecture_history_id `
	extract_sentence += where_phase

	fmt.Println(extract_sentence)
	rows, err := db.Query(extract_sentence)
	defer rows.Close()

	resJson := []ReturnReload{}
	if err != nil {
		resJson = append(resJson, ReturnReload{})
		resJson[0].Subject_name = ""
		resJson[0].Http_status = http.StatusCreated
	} else {
		for rows.Next() {
			res := ReturnReload{}
			rows.Scan(&res.Subject_name)
			res.Http_status = http.StatusCreated
			fmt.Println(res)
			resJson = append(resJson, res)
		}
	}

	return c.JSON(http.StatusCreated, resJson)
}
