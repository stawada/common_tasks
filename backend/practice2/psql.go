package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

/* 出欠情報用のjson */
type Attend struct {
	student_id   string `json:"student_id"`
	student_name string `json:"student_name"`
	hashed_pw    string `json:"hashed_pw"`
}

var db *sql.DB

func init() {
	_, err := sql.Open("postgres", "user=stawada dbname=sample")
	if err != nil {
		panic("Not found Database.")
	}
}

func GetFromPath(c echo.Context) error {
	no := c.Param("no")
	res := Attend{}

	if err := db.QueryRow("select student_id,student_name,hashed_pw from sample", no).Scan(&res.student_id, &res.student_name, &res.hashed_pw); err != nil {
		panic("Did not read!")
	}

	return c.JSON(http.StatusCreated, res)
}
