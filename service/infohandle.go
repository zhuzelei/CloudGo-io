package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func InfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		//输出匿名结构并json序列化输出
		formatter.HTML(w, http.StatusOK, "table", struct {
			Name  string
			Phone string
		}{Name: req.Form["name"][0], Phone: req.Form["phone"][0]})
	}
}
