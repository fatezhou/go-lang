package server

import (
	"sync"
	stlTemplate "text/template"
)


var templateCache *sync.Map

type Template struct{
}

func init(){
	templateCache = &sync.Map{}
}

func GetTemplate(strPath string)*stlTemplate.Template{
	strPath = "html/" + strPath
	cache, ok := templateCache.Load(strPath)
	if ok {
		data, ok := cache.(*stlTemplate.Template)
		if ok{
			return data
		}else{
			return nil
		}
	}else{
		t, err := stlTemplate.ParseFiles(strPath)
		if err != nil{
			return nil
		}else{
			templateCache.Store(strPath, t)
			return t
		}
	}
}
