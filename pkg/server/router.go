package server

import (
	"git.jiaxianghudong.com/go/logs"
	"net/http"
	"strings"
)

func GetPath(strUrl string) string {
	if len(strUrl) > 0 && []byte(strUrl)[0] == '/' {
		return strings.ToLower(strUrl[1:])
	} else {
		return ""
	}
}

func IsApi(strUrl string) bool {
	arr := strings.Split(strUrl, "/")
	if len(arr) == 0 {
		return true
	}

	lastWord := arr[len(arr)-1]
	for _, word := range lastWord {
		if word == '.' {
			return false
		}
	}
	return true
}

func Router(writer http.ResponseWriter, request *http.Request) {
	strUrl := strings.ToLower(request.URL.Path)
	strPath := GetPath(strUrl)
	logs.Infof("path:%s", strPath)
	result := false
	defer func() {
		if result == false {
			writer.WriteHeader(http.StatusServiceUnavailable)
			writer.Write([]byte(""))
		}
	}()
	hc := GetHttpContext()
	hc.Fill(writer, request)
	result = Handle(hc)
}
