package utils

import "github.com/fatezhou/go-lang"

var t = zoyee_go_lang.Timer{}

func init(){
	t.Init(500, 0)
}

func SetTimer(deskId int32, timerId int32, duration int32, timerProc func(deskId int32, timerId int32, extData interface{}), extData interface{}){
	t.SetTimer(deskId, timerId, duration, timerProc, extData)
}

func KillTimer(deskId int32, timerId int32){
	t.KillTimer(deskId, timerId)
}
