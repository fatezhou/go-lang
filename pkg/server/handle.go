package server

import (
	"sync"
)

type HttpHandle func(hc *HttpContext) bool

type handle struct {
	m       sync.Mutex
	handles map[string]HttpHandle
}

var handleInstance handle

func getHandle(strUrl string) HttpHandle {
	handleInstance.m.Lock()
	defer handleInstance.m.Unlock()
	if v, ok := handleInstance.handles[strUrl]; ok {
		return v
	} else {
		return nil
	}
}

func init() {
	handleInstance = handle{
		m:       sync.Mutex{},
		handles: make(map[string]HttpHandle),
	}
}

func AddHandle(strUrl string, iHandle HttpHandle) {
	handleInstance.m.Lock()
	defer handleInstance.m.Unlock()
	handleInstance.handles[strUrl] = iHandle
}

func Handle(hc *HttpContext) (res bool) {
	res = false
	defer PutHttpContext(hc)
	if handleFunc := getHandle(hc.Path); handleFunc != nil {
		res = handleFunc(hc)
	}
	if res == false {
		return
	}
	hc.Resp()
	return
}
