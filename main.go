package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	var indexdata IndexData
	indexdata.Title = "锦先生的个人博客项目"
	indexdata.Desc = "这是主页"
	t := template.New("index.html")
	//1.拿到当前的路径
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	t, err = t.ParseFiles(path + "/template/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, indexdata) //传入index.html页面中
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//处理请求，返回响应
	http.HandleFunc("/index.html", Index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
