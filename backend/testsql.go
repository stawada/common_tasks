// testsql.go
package main

import (
	"database/sql"
	"fmt"

	_"github.com/lib/pq"
)

type Sale struct {
	Id      string
	OrderId string
}

const (
	HOST = "localhost"    //ホスト名
	DATABASE = "testdb"    //データベースの名前
	USER = "yuusyamada"    //ユーザーネーム
	PASSWORD = "Pasona123"
)

func checkError(err error){
	if err != nil {
		panic(err)
	}
}

func main(){
	var connectionString string = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", HOST, USER, DATABASE)  //sslmode=disableにてSSLモードを無効

	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	    fmt.Println("Finished dropping table (if existed)")

    // Create table.
    _, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")  //表の作成
    checkError(err)
    fmt.Println("Finished creating table")

    // Insert some data into table.
    sql_statement := "INSERT INTO inventory (name, quantity) VALUES ($1, $2);"
    _, err = db.Exec(sql_statement, "banana", 150)
    checkError(err)
    _, err = db.Exec(sql_statement, "orange", 154)
    checkError(err)
    _, err = db.Exec(sql_statement, "apple", 100)
    checkError(err)
    fmt.Println("Inserted 3 rows of data")
}
