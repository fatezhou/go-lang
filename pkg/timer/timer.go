package timer

import (
	"fmt"
	"github.com/fatezhou/go-lang"
	"sync"
	"time"
)

type TimerInfo struct{
	TimerProc func(deskId int32, timerId int32, extData interface{})
	TimerId int32
	Duration int32
	DeskId int32
	ExtData interface{}
	stopChan chan int
	Ticker *time.Ticker
}

type TimerMap struct{
	Map *sync.Map
}

func (t* TimerInfo) stop(){
	close(t.stopChan)
	t.Ticker.Stop()
}

var timerMap TimerMap

func init(){
	timerMap.Map = &sync.Map{}
}

func SetTimer(deskId int32, timerId int32, duration int32, timerProc func(deskId int32, timerId int32, extData interface{}), extData interface{}){
	strTimeId := fmt.Sprintf("%d_%d", deskId, timerId)
	zoyee_go_lang.Info("TimeId:%s" ,strTimeId)
	if v, ok := timerMap.Map.Load(strTimeId); ok{
		if t ,ok := v.(*TimerInfo) ;ok {
			t.stop()
		}
		//KillTimer(deskId, timerId)
	}
	timer := &TimerInfo{
		TimerProc:timerProc,
		TimerId:timerId,
		DeskId:deskId,
		Duration:duration,
		ExtData:extData,
		Ticker: time.NewTicker(time.Duration(duration) * time.Millisecond),
		stopChan:make(chan int),
	}
	timerMap.Map.Store(strTimeId ,timer)
	go func(timer* TimerInfo) {
		for{
			select{
			case <-timer.Ticker.C:
				timer.TimerProc(timer.DeskId, timer.TimerId, timer.ExtData)
			case _ ,ok := <-timer.stopChan:
				if ok {
					fmt.Println("timer.stopChan :" ,timer.TimerId)
					timer.stop()
					fmt.Println("timer.stopChan end :" ,timer.TimerId)
				}
				return
			}
		}
	}(timer)
}

func KillTimer(deskId int32, timerId int32){
	go func() {
		strTimeId := fmt.Sprintf("%d_%d", deskId, timerId)
		if v, ok := timerMap.Map.Load(strTimeId); ok{
			timer , _ := v.(*TimerInfo)
			timer.stopChan<-1
			timerMap.Map.Delete(strTimeId)
		}else{
			zoyee_go_lang.Infof("can not find timer, kill fail, deskId:[%d], timerId:[%d]", deskId, timerId)
		}
	}()
}