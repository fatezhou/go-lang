package http

import (
	"fmt"
	"net/http"
	"strings"
)

type CmdHandler func(data map[string]string)(strResponseText string)

type HttpServer struct{
	router map[string]CmdHandler
}

func (httpServer *HttpServer)ServeHTTP(writer http.ResponseWriter, request *http.Request){
	strCmd := strings.ToLower(request.URL.Path)
	if handle, ok := httpServer.router[strCmd]; ok{
		mapData := make(map[string]string)
		value := request.URL.Query()
		for k, v := range value{
			if len(v) > 0{
				mapData[k] = v[0]
			}
		}
		response := handle(mapData)
		strJson := fmt.Sprintf(`
{
	"code":0, 
	"data":%s
}`		, response)
		writer.Write([]byte(strJson))
	}else{
		writer.Write([]byte(`{"code":-1, "data":{"text":"type http://ip:port/help to get cmds"}}`))
	}
}

func (httpServer *HttpServer)AddHandle(cmd string, handle CmdHandler){
	//     cmd:  /cmd/xxxxx
	cmd = strings.ToLower(cmd)
	if httpServer.router == nil{
		httpServer.router = make(map[string]CmdHandler)
	}
	httpServer.router[cmd] = handle
}

func (httpServer *HttpServer)Run(port int){
	strHttpIp := fmt.Sprintf("0.0.0.0:%d", port)
	http.ListenAndServe(strHttpIp, httpServer)
}

func (httpServer *HttpServer)RunAync(port int){
	go func(){
		httpServer.Run(port)
	}()
}