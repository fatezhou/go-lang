package server

import (
	"buyu-web/pkg/tools"
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var pool sync.Pool

func init() {
	pool = sync.Pool{}
	pool.New = func() interface{} {
		return NewHC()
	}
}

type HttpContext struct {
	Ctx          context.Context
	Method       string
	Path         string
	Params       map[string]string
	Header       map[string]string
	Cookie       map[string]string
	Body         []byte
	ResponseBody []byte
	Response     http.ResponseWriter
	Code         int
}

func GetHttpContext() (hc *HttpContext) {
	return pool.Get().(*HttpContext)
}

func PutHttpContext(hc *HttpContext) {
	hc.Reset()
	pool.Put(hc)
}

func (hc *HttpContext) Fill(writer http.ResponseWriter, request *http.Request) {
	hc.Response = writer
	hc.Method = request.Method
	for k, v := range request.Header {
		hc.Header[k] = v[0]
	}
	for _, v := range request.Cookies() {
		hc.Cookie[v.Name] = v.Value
	}
	for k, v := range request.URL.Query() {
		hc.Params[k] = v[0]
	}
	var err error
	hc.Body, err = ioutil.ReadAll(request.Body)
	if err == nil {
		request.Body.Close()
	}
	hc.Code = http.StatusOK
	hc.Path = GetPath(strings.ToLower(request.URL.Path))
}

func NewHC() *HttpContext {
	hc := &HttpContext{}
	hc.Params = make(map[string]string)
	hc.Cookie = make(map[string]string)
	hc.Header = make(map[string]string)
	hc.Body = make([]byte, 0)
	hc.ResponseBody = make([]byte, 0)
	return hc
}

func (hc *HttpContext) Reset() {
	hc.Params = make(map[string]string)
	hc.Cookie = make(map[string]string)
	hc.Header = make(map[string]string)
	hc.Body = make([]byte, 0)
	hc.ResponseBody = make([]byte, 0)
}

func (hc *HttpContext) Resp() {
	if len(hc.Cookie) > 0 {
		for k, v := range hc.Cookie {
			http.SetCookie(hc.Response, &http.Cookie{Name: k, Value: v})
		}
	}
	if len(hc.Header) > 0 {
		for k, v := range hc.Header {
			hc.Response.Header().Set(k, v)
		}
	}
	hc.Response.WriteHeader(hc.Code)
	hc.Response.Write(hc.ResponseBody)
}

func (hc *HttpContext) GetCookie(key string) string {
	if v, ok := hc.Cookie[key]; ok {
		return v
	} else {
		return ""
	}
}

func (hc *HttpContext) GetHeader(key string) string {
	if v, ok := hc.Header[key]; ok {
		return v
	} else {
		return ""
	}
}

func (hc *HttpContext) GetParam(key, defaultValue string) string {
	if v, ok := hc.Params[key]; ok {
		return v
	} else {
		return defaultValue
	}
}

func (hc *HttpContext) GetParamToInt32(key string, defaultValue int32) int32 {
	if v, ok := hc.Params[key]; ok {
		return tools.Str2Int32(v)
	}
	return defaultValue
}

func (hc *HttpContext) CheckParam(key string, isExist bool) bool {
	_, ok := hc.Params[key]
	if isExist && ok {
		return true
	}
	if isExist && ok == false {
		return false
	}
	if isExist == false && ok == false {
		return true
	}
	if isExist == false && ok == true {
		return true
	}
	return false
}

func (hc *HttpContext) CheckParams(keys []string, ifExist bool) (res bool, key string) {
	for _, v := range keys {
		if hc.CheckParam(v, ifExist) == false {
			return false, v
		}
	}
	return true, ""
}
