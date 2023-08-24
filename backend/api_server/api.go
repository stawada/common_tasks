package main

import (
	"database/sql"
	"fmt"

	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "host=db user=postgres password=Pasona123 sslmode=disable")
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
	if err := db.QueryRow(select_sentence).Scan(&checkJson.Student_id, &checkJson.Hashed_password); err != nil {
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
	Now_time    int    `json:"now_time"`
	Student_id  string `json:"student_id"`
	Subject_id  string `json:"subject_id"`
}

type ReturnAttendInfo struct {
	Check_flag  int `json:"check_flag"`
	Http_status int `json:"http_status"`
}

func PostAttend(c echo.Context) error {
	post := new(ReceiveAttendInfo)
	err := c.Bind(post)
	if err != nil {
		return err
	}

	check_flag := 1 // 1->success, 0->denied
	check_sentence := fmt.Sprintf(`SELECT lecture_catalog.lecture_name FROM lecture_catalog 
						INNER JOIN lecture_history ON lecture_catalog.lecture_catalog_id=lecture_history.lecture_catalog_id 
						INNER JOIN attendance_information ON lecture_history.lecture_history_id=attendance_information.lecture_history_id `)
	if post.Attend_flag == -1 {
		check_sentence += fmt.Sprintf("WHERE attendance_information.student_id='%s';", post.Student_id)
	} else {
		check_sentence += fmt.Sprintf("WHERE lecture_history.lecture_date_and_time-600<=%d AND lecture_history.lecture_date_and_time+600>=%d AND attendance_information.student_id='%s';", post.Now_time, post.Now_time, post.Student_id)
	}

	var subject_name string
	resJson := ReturnAttendInfo{}
	if err := db.QueryRow(check_sentence).Scan(&subject_name); err != nil || subject_name == "" {
		check_flag = 1 // エラー時はフラグを変更
	} else {
			update_sentence := fmt.Sprintf(`UPDATE attendance_information AS atnd 
											SET attendance_flag=%d 
											WHERE attendance_information_id=(
												SELECT attendance_information_id FROM attendance_information AS atnd 
												INNER JOIN lecture_history AS hist ON hist.lecture_history_id=atnd.lecture_history_id 
												INNER JOIN lecture_catalog AS ctlg ON ctlg.lecture_catalog_id=hist.lecture_catalog_id 
												WHERE atnd.student_id='%s' AND ctlg.lecture_catalog_id='%s');`, post.Attend_flag, post.Student_id, post.Subject_id)
		if _, err := db.Exec(update_sentence); err != nil {
			check_flag = 0
		}
	}

	resJson.Check_flag = check_flag
	resJson.Http_status = http.StatusCreated
	return c.JSON(http.StatusCreated, resJson)
}

/* 講義情報送信用API */
type ReceiveReload struct {
	Student_id string `json:"student_id"`
	Now_time   int    `json:"now_time"`
	Attend_flag int `json:"attend_flag"`
}

type ReturnReload struct {
	Subject_name string `json:"subject_name"`
	Subject_id   string `json:"subject_id"`
	Subject_time int    `json:"subject_time"` // UnixTime
	Http_status  int    `json:"http_status"`
}

func PostReload(c echo.Context) error {
	post := new(ReceiveReload)
	err := c.Bind(post)
	if err != nil {
		return err
	}

	var where_phase string
	if post.Attend_flag == -1 {
		where_phase = fmt.Sprintf("WHERE attendance_information.student_id='%s' AND lecture_history.lecture_date_and_time>=%d;", post.Student_id, post.Now_time)
	} else {
		where_phase = fmt.Sprintf("WHERE lecture_history.lecture_date_and_time-600<=%d AND lecture_history.lecture_date_and_time+600>=%d AND attendance_information.student_id='%s';", post.Now_time, post.Now_time, post.Student_id)
	}
	extract_sentence := `SELECT lecture_catalog.lecture_name, lecture_catalog.lecture_catalog_id, lecture_history.lecture_date_and_time, attendance_information.attendance_flag FROM lecture_catalog 
						INNER JOIN lecture_history ON lecture_catalog.lecture_catalog_id=lecture_history.lecture_catalog_id 
						INNER JOIN attendance_information ON lecture_history.lecture_history_id=attendance_information.lecture_history_id `
	extract_sentence += where_phase

	fmt.Println(extract_sentence)
	rows, err := db.Query(extract_sentence)
	defer rows.Close()

	resJson := make([]ReturnReload, 0)
	if err != nil {
		resJson = append(resJson, ReturnReload{})
		resJson[0].Subject_name = ""
		resJson[0].Subject_id = ""
		resJson[0].Subject_time = 0
		resJson[0].Http_status = http.StatusCreated
	} else {
		for rows.Next() {
			var attend_flag int
			res := ReturnReload{}
			rows.Scan(&res.Subject_name, &res.Subject_id, &res.Subject_time, &attend_flag)
			// 出席済み or 事前欠席済みの場合はreturnのデータに含めない
			if attend_flag != 0 {
				continue
			}
			res.Http_status = http.StatusCreated
			resJson = append(resJson, res)
		}
	}

	return c.JSON(http.StatusCreated, resJson)
}
