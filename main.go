package main

import (
	"GoBlog/config"
	"GoBlog/models"
	"GoBlog/router"
	"html/template"
	"log"
	"net/http"
	"time"
)

func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	t := template.New("index.html")
	//1.拿到当前的路径
	path := config.Cfg.System.CurrentDir
	//访问博客首页模板的时候，因为有多个模板的嵌套，解析文件的时候，需要将其涉及到的所有模板都进行解析
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t.Funcs(template.FuncMap{
		"isODD":       IsODD,
		"getNextName": GetNextName,
		"date":        Date,
	})
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)
	if err != nil {
		log.Println(err)
	}
	//页面上涉及到的所有数据，都要有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var homeresponse = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	t.Execute(w, homeresponse) //传入index.html页面中
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//处理请求，返回响应
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
