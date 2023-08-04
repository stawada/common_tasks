// demo_dbを操作する用のコード
package main

import (
	"database/sql" //DB操作
	"fmt"
	"log"

	_"github.com/lib/pq"
)

type App struct {
	Student_id	string
	Student_name string
	Hashed_password string
	Attendance_information_id string
	Attendance_flag int
	Lecture_catalog_id string
	Lecture_id string
	Lecture_name string
	Lecture_history_id string
	lecture_data_and_time int
}

const ( //接続するDBの情報
	HOST = "localhost"    //ホスト名
	DATABASE = "demo_db"    //データベースの名前
	USER = "yuusyamada"    //ユーザーネーム
	PASSWORD = "Pasona123"
)

func checkError(err error){ //エラーの確認
	if err != nil {
		panic(err)
	}
}

func attendance_update(){ //出欠席
	var connectionString string = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", HOST, USER, DATABASE)  //sslmode=disableにてSSLモードを無効

	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)

	sql_statement := `UPDATE attendance_information 
					  SET attendance_flag = %d
					  WHERE student_id='%s' AND lecture_history_id='%s';` //update文の実行(フラグ管理)
	_, err = db.Exec(sql_statement)
	checkError(err)
}

func main(){//出席可能講義
	var connectionString string = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", HOST, USER, DATABASE)  //sslmode=disableにてSSLモードを無効

	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)

	sql_statement := `SELECT lecture_catalog.lecture_name
					  FROM((attendance_information
						INNER JOIN lecture_history ON attendance_information.lecture_history_id = lecture_history.lecture_history_id)
						INNER JOIN lecture_catalog ON lecture_history.lecture_catalog_id = lecture_history.lecture_catalog_id)
				 	 WHERE student_id = '1000000000' and lecture_history.lecture_data_and_time-5 <= 1000 and lecture_history.lecture_data_and_time+5 >= 1000;` //SELECT文の実行(出席可能講義の抽出)
	rows, err := db.Query(sql_statement)
	checkError(err)
	if err != nil{
		log.Println(err)
	}
	defer rows.Close()

	var u App

	for rows.Next(){
		err := rows.Scan(&u.Lecture_name)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Printf("%s\n", u.Lecture_name)
	}
	err = rows.Err()
	if err != nil{
		log.Fatal(err)
	}
}
