package main

import (
	"fmt"
	"sync"
	"time"
	ul "zoyee-tool/internal/pkg/utils"
)

var BuildTime = ""
var Debug = ""

func GetThisMonday() int32 {
	now := time.Now()
	return int32(now.YearDay()) - int32(now.Weekday()) + 1 //取最近一个周一是这一年内的第几天
}

func Today() int{
	now := time.Now()
	return now.YearDay()
}
var wg = sync.WaitGroup{}
var i = 0
var t ul.Timer
func OnTimer(id1, id2 int32, extData interface{}){
	fmt.Printf("[%d][%d][%d]%+v\n", id1, id2, time.Now().Unix(), extData)
	if id1 == 100{
		wg.Done()
	}
	i++
	if i >= 5{
		t.KillTimer(1, 1)
	}

}

func main(){
	wg.Add(1)

	t.Init(500)
	t.SetTimer(1, 1, 1500, OnTimer, nil)
	t.SetTimer(2, 2, 1500, OnTimer, nil)
	wg.Wait()
}
