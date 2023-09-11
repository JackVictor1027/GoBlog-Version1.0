package router

import "net/http"

func Router() {

	http.HandleFunc("/index.html", Index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
