package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"time"
)

var (
	baseUrl = "https://api.openai.com/v1/chat/completions"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/chat/completions", HandleProxy)
	if err := http.ListenAndServe(":80", mux); err != nil {
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
	dumpreq, _ := httputil.DumpRequest(r, true)
	newreq, _ := http.ReadRequest(bufio.NewReader(bytes.NewBuffer(dumpreq)))
	req, _ := http.NewRequest(newreq.Method, baseUrl, newreq.Body)
	for k, v := range r.Header {
		req.Header[k] = v
	}
	rsp, err := client.Do(newreq)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}
	data, _ := httputil.DumpResponse(rsp, true)
	w.Write(data)
}
