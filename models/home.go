package models

import "GoBlog/config"

type HomeResponse struct { //不用HomeData而改用HomeResponse
	config.Viewer
	Categorys []Category
	Posts     []PostMore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}
