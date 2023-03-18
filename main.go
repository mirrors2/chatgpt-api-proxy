package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	baseUrl = "https://api.openai.com"
)

func main() {
	router := mux.NewRouter()
	// 路由转发
	router.HandleFunc("/v1/chat/completions", HandleProxy)

	// 启动代理服务器
	fmt.Println("API proxy server is listening on port 8080")
	if err := http.ListenAndServe(":80", router); err != nil {
		panic(err)
	}
}

func HandleProxy(w http.ResponseWriter, r *http.Request) {
	client := http.DefaultClient
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
	client.Transport = tr

	// 创建 API 请求
	req, err := http.NewRequest(r.Method, baseUrl+r.URL.Path, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header = r.Header

	rsp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rsp.Body.Close()

	// 复制 API 响应头部
	for name, values := range rsp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	// 返回 API 响应主体
	w.WriteHeader(rsp.StatusCode)
	if _, err := io.Copy(w, rsp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
