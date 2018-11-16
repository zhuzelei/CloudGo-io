package service

//使用文件服务

import (
	"net/http"
	"os"

	
	"github.com/codegangsta/negroni" 
	"github.com/gorilla/mux"         
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic() 
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx) 
	return n
}

//
func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}

	mx.HandleFunc("/unknown", UnknownHandler(formatter))
	mx.HandleFunc("/table", InfoHandler(formatter))
	mx.HandleFunc("/js", JsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))

	/*
		将 path 以 “/” 前缀的 URL 都定位到 webRoot + "/assets/" 为虚拟根目录的文件系统。
		http.Dir 是类型。将字符串转为 http.Dir 类型，这个类型实现了 FileSystem 接口。（Dir 不是函数）
		http.FileServer() 是函数，返回 Handler 接口，该接口处理 http 请求，访问 root 的文件请求。
		mx.PathPrefix 添加前缀路径路由。
	*/
}
