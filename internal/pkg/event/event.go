package event

import "sync"

type eventInfo struct {
	eventName string
	eventParam []interface{}
}

type EventListener struct{
	name string
	events *sync.Map
	eventMQ chan *eventInfo
	infoPool *sync.Pool
}

var Event EventListener

func init(){
	Event.New("global_event")
}

func (e *EventListener)New(name string){
	e.events = &sync.Map{}
	e.infoPool = &sync.Pool{
		New: func() interface{} {
			return makeEventInfo()
		},
	}
	e.eventMQ = make(chan *eventInfo)
	e.name = name
	go e.Disptch()
}

func makeEventInfo()*eventInfo {
	eventInfo := &eventInfo{}
	eventInfo.eventParam = make([]interface{}, 0)
	return eventInfo
}

func (ev *eventInfo)reset(){
	ev.eventParam = make([]interface{}, 0)
	ev.eventName = ""
}

func (ev *eventInfo)push(args ...interface{}){
	for _, v := range args{
		ev.eventParam = append(ev.eventParam, v)
	}
}

func (e *EventListener)getEventInfo(eventName string)*eventInfo{
	ev := e.infoPool.Get().(*eventInfo)
	ev.eventName = eventName
	return ev
}

func (e *EventListener)putEventInfo(info *eventInfo){
	info.reset()
	e.infoPool.Put(info)
}

func (e *EventListener)Emit(eventName string, args ...interface{}){
	eventInfo := e.getEventInfo(eventName)
	eventInfo.push(args...)
	e.eventMQ <- eventInfo
}

func (e *EventListener)On(eventName string, pFn func(args ...interface{})){
	e.events.Store(eventName, pFn)
}

func (e *EventListener)Disptch(){
	for{
		data := <-e.eventMQ
		value, err := e.events.Load(data.eventName)
		if err != true{
			e.putEventInfo(data)
			continue
		}
		pfn := value.(func(args ...interface{}))
		pfn(data.eventParam...)
		e.putEventInfo(data)
	}
}

func Emit(eventName string, args ...interface{}){
	Event.Emit(eventName, args...)
}

func On(eventName string, pFn func(args ...interface{})){
	Event.On(eventName, pFn)
}