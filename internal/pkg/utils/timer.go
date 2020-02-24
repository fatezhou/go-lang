package utils

import (
	"sync"
	"time"
)

type TimerCallbackFunc func (deskId, playerId int32, extData interface{})

type TimerInfo struct{
	Data interface{}
	Duration int32
	NextAwakeTimestamp int64
	TimerId int32
	TimerId2 int32
	callback TimerCallbackFunc
}

type Timer struct {
	timers *sync.Map
	core *time.Ticker
	init bool
	seq uint64
	callbackChan chan *TimerInfo
	lock sync.Mutex
}

func (t *Timer)GetSeq() uint64{
	//return atomic.AddUint64(&t.seq, 1)
	t.lock.Lock()
	defer t.lock.Unlock()
	t.seq++
	return t.seq
}

func (t *Timer)Init(interval int32){
	if t.init{
		return
	}
	if interval <= 0{
		interval = 1000
	}
	t.core = time.NewTicker(time.Duration(interval) * time.Millisecond)
	t.timers = &sync.Map{}
	t.init = true
	t.seq = 0
	t.callbackChan = make(chan *TimerInfo, 0xFFFF)
	go t.Callback()
	go t.Loop()
}

func (t *Timer)KillTimer(id1 int32, id2 int32){
	t.timers.Range(func(key, value interface{}) bool {
		timerInfo := value.(*TimerInfo)
		if timerInfo.TimerId == id1 && timerInfo.TimerId2 == id2{
			t.timers.Delete(key)
			return false
		}
		return true
	})
}

func (t *Timer)GetMs()int64{
	second := time.Now().UnixNano()/ 1e6
	return second
}

func (t *Timer)SetTimer(id1 int32, id2 int32, duration int32, timerProc func (deskId, playerId int32, extData interface{}), extData interface{}){
	timerInfo := &TimerInfo{
		Data: extData,
		Duration: duration,
		NextAwakeTimestamp: t.GetMs() + int64(duration),
		TimerId: id1,
		TimerId2: id2,
		callback: timerProc,
	}
	t.timers.Store(t.GetSeq(), timerInfo)
}

func (t *Timer)Callback(){
	for{
		timerInfo := <-t.callbackChan
		if timerInfo != nil{
			timerInfo.callback(timerInfo.TimerId, timerInfo.TimerId2, timerInfo.Data)
		}
	}
}

func (t *Timer)Loop(){
	for{
		ms := t.GetMs()
		select{
		case <-t.core.C:{
				t.timers.Range(func(_, value interface{}) bool {
					timerInfo := value.(*TimerInfo)
					//fmt.Printf("s:%d, %+v\n", second, timerInfo)
					if timerInfo.NextAwakeTimestamp <= ms{
						//触发定时
						timerInfo.NextAwakeTimestamp += int64(timerInfo.Duration)
						t.callbackChan <- timerInfo
					}
					return true
				})
			}
		}
	}
}

