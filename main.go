package main

import (
	"database/sql"
	"html"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"./config"
)

type Comment struct {
	Name string
	Body string
}

func main() {
	// static files(CSS and JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", pageHandler)     // ページの処理
	http.HandleFunc("/post", postHandler) // POSTの処理
	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

// ページの処理
func pageHandler(w http.ResponseWriter, r *http.Request) {
	msg := selectComment("SELECT * FROM test1 LIMIT 64")
	tpl, err1 := template.ParseFiles("view/index.html.tpl")
	if err1 != nil {
		panic(err1)
	}
	err2 := tpl.ExecuteTemplate(w, "index.html.tpl", msg)
	if err2 != nil {
		panic(err2)
	}
}

// POSTの処理
func postHandler(w http.ResponseWriter, r *http.Request) {
	name := html.EscapeString(r.FormValue("Name"))
	body := html.EscapeString(r.FormValue("Body"))
	log.Printf("Name=%s Body=%s", name, body)
	insertDB("INSERT INTO test1(name,body) VALUES(\"" + name + "\",\"" + body + "\")")
	http.Redirect(w, r, "/", 301) // 元のパスにRedirectする
}

// DBから読み込み
func selectComment(q string) [64]Comment {
	var data Comment
	var msg [64]Comment

	db, err := sql.Open("mysql", config.DB.User+":"+config.DB.Password+"@/"+config.DB.Name)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Select文発行
	rows, err := db.Query(q)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// 値の取得
	i := 0
	for rows.Next() {
		if err := rows.Scan(&(data.Name), &(data.Body)); err != nil {
			panic(err.Error())
		}
		msg[i] = data
		i++
	}
	return msg
}

// DBへの書き込み
func insertDB(q string) {
	db, err := sql.Open("mysql", config.DB.User+":"+config.DB.Password+"@/"+config.DB.Name)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Insert文発行
	result, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}
	// 影響を与えた行数を返す
	n, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	log.Printf("inserted:%d\n", n)
}
