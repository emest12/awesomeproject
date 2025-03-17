package http

import (
	"net/http"
	"testing"
)

func TestHttpHandler1(t *testing.T) {
	// 把一个 HTTP 处理函数注册到指定的路径上
	http.HandleFunc("/gwhtest", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// 启动 HTTP 服务器，监听指定的端口
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

	// 另起终端 curl -v http://127.0.0.1:8080/gwhtest 来验证
}
