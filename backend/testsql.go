// testsql.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_"github.com/lib/pq"
)

var (
	book_number *int
	book_name   *string
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
/*   _, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")  //表の作成
    checkError(err)
    fmt.Println("Finished creating table")
*/
    // Insert some data into table.
    sql_statement := "SELECT * FROM books;"
	rows, err := db.Query(sql_statement)
	if err != nil{
		log.Println(err)
	}
	defer rows.Close()

	var u Sale

	for rows.Next(){
		err := rows.Scan(&u.Id, &u.OrderId)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Printf("ID:%s| Name:%s\n", u.Id, u.OrderId)
	}
	err = rows.Err()
	if err != nil{
		log.Fatal(err)
	}
	//checkError(err)

	//fmt.Print(sample)
	//fmt.Println("update ぐりとぐら => てつとゆめ")
}

