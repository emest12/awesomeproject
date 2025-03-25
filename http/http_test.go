package http

import (
	"golang.org/x/net/context"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestHttpHandler1(t *testing.T) {
	// 把一个 HTTP 处理函数注册到指定的路径上
	http.HandleFunc("/gwhtest", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})

	// 启动 HTTP 服务器，监听指定的端口
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

	// 另起终端 curl -v http://127.0.0.1:8080/gwhtest 来验证
	// 或者通过客户端验证
}

func TestHttpClient(t *testing.T) {
	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunc()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://127.0.0.1:8080/gwhtest", nil)
	if err != nil {
		t.Error(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}
		t.Log("body:", string(body))
	}
}
