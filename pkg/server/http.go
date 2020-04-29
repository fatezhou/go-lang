package server

import (
	"fmt"
	"git.jiaxianghudong.com/go/logs"
	"net/http"
)

type HttpServer struct {
}

func (hs *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	Router(writer, request)
}

func (hs *HttpServer) Run(port int) {
	strServerIpAddr := fmt.Sprintf("0.0.0.0:%d", port)
	logs.Infof("Server run in [%s]", strServerIpAddr)
	http.ListenAndServe(strServerIpAddr, hs)
}
