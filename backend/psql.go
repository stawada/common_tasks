// demo_dbを操作する用のコード
package main

import (
	//"encoding/json" //json関連
	"database/sql" //DB操作
	"fmt"
	//"io/ioutil" //ファイル読み込み

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

/*func readJson(){ //jsonの読み込み
	s, _ := ioutil.ReadFile("sample_1.json")
	var jsondata map[string]interface{}
	json.Unmarshal([]byte(s), &jsondata)
	fmt.Println(jsondata)

	return
}*/

func main(){ //DBへの書き込み
	var connectionString string = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", HOST, USER, DATABASE)  //sslmode=disableにてSSLモードを無効

	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	    //fmt.Println("Finished dropping table (if existed)")

	//readJson() //jsonファイルの読み込み
    sql_statement := "UPDATE attendance_information SET attendance_flag = %d WHERE student_id='%s' AND lecture_history_id='%s';", //update文の実行(フラグ管理)
	_, err = db.Exec(sql_statement)
	checkError(err)

}