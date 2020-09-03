package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"time"
)

type Class struct {
	ID     int
	Date   time.Time
	Time   string
	Mentor string
}

//temp1は1つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	temp1    *template.Template
	str      []Class
}

//ServeHTTPはHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.temp1 = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.temp1.Execute(w, t.str)
}

var db *sql.DB

var err error

func main() {
	//mysqlへ接続。ドライバ名（mysql）と、ユーザー名・データソース(ここではgosample)を指定。
	db, err = sql.Open("mysql", "root:root@tcp(godockerDB)/mysql?parseTime=true")
	log.Println("Connected to mysql.")

	//接続でエラーが発生した場合の処理
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	rows, err := db.Query(`SELECT * FROM class;`)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	var data []Class
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var class Class //構造体Person型の変数personを定義
		err = rows.Scan(&class.ID, &class.Date, &class.Time, &class.Mentor)

		if err != nil {
			panic(err.Error())
		}
		data = append(data, class)
	}

	// ルート
	http.Handle("/", &templateHandler{filename: "chat.html", str: data})

	//Webサーバーを開始します
	if err = http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
