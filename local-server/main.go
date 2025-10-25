package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

const baseMessage = "Hello World"

func handler2(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.FormValue("name")
	if name == "" {
		fmt.Fprint(w, baseMessage)
		return
	} else {
		user := User{
			Name: name,
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(user)
	}

	fmt.Fprintf(w, "%s, %s", baseMessage, name)
}

func main() {
	//ディレクトリを指定する
	fs := http.FileServer(http.Dir("static"))
	//ルーティング設定。"/"というアクセスがきたらstaticディレクトリのコンテンツを表示させる
	http.Handle("/", fs)

	http.HandleFunc("/test", handler)
	http.HandleFunc("/test2", handler2)

	log.Println("Listening...")
	// 3000ポートでサーバーを立ち上げる
	http.ListenAndServe(":8080", nil)
}
