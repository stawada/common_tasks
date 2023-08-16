package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func calc_timeLimit() time.Time {
	now_time := time.Now()
	y := now_time.Year()
	m := now_time.Month()
	d := now_time.AddDate(0, 0, 1).Day()
	all_day := time.Date(y, m, d, 0, 0, 0, 0, now_time.Location())
	return all_day
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:    "token",
		Value:   "sample",
		Path:    "/",
		Expires: calc_timeLimit(),
		// Secure: os.Getenv("COOKIE_SECURE"),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		panic(err)
	}

	if token.Value == "sample" {
		fmt.Println(token.Value)
		tmpl := template.Must(template.ParseFiles("./index.html"))
		tmpl.Execute(w, token)
	}
}

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/cookie", readCookie)

	if err := http.ListenAndServe(":1323", nil); err != nil {
		panic(err)
	}
}
