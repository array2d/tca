package apiserver

import "github.com/gorilla/mux"

func NewRouter() (root *mux.Router) {
	router := mux.NewRouter()

	//当请求的URL路径以斜线结尾时，mux会自动修正为不以斜线结尾的形式进行匹配
	router.StrictSlash(true)
	//router.Use(middleware.Logging)
	return router
}
