package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

var (
	baseUrl = "https://api.openai.com"
)

func main() {
	router := http.NewServeMux()
	// 路由转发
	router.HandleFunc("/", HandleProxy)

	// 启动代理服务器
	fmt.Println("API proxy server is listening on port 80")
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
	head := map[string]string{
		"Cache-Control":                    "no-store",
		"access-control-allow-origin":      "*",
		"access-control-allow-credentials": "true",
	}
	for k, v := range head {
		if _, ok := rsp.Header[k]; !ok {
			w.Header().Set(k, v)
		}
	}
	rsp.Header.Del("content-security-policy")
	rsp.Header.Del("content-security-policy-report-only")
	rsp.Header.Del("clear-site-data")

	// 返回 API 响应主体
	w.WriteHeader(rsp.StatusCode)
	if _, err := io.Copy(w, rsp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
