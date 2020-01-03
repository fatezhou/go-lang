package lua

import(
	"bufio"
	"container/list"
	"errors"
	"fmt"
	glua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"os"
	"runtime/debug"


	//"os"
	"sync"
)

var stdout func(strFmt string, value ...interface{})

func init(){
	stdout = nil
}

func SetStdOut(funcPtr func(strFmt string, value ...interface{})){
	stdout = funcPtr
}

func out(strFmt string, value ...interface{}){
	if stdout != nil{
		stdout(strFmt, value...)
	}else{
		fmt.Printf(strFmt, value...)
	}
}

type Lua struct {
	luaPath  string
	luaProto *glua.FunctionProto
	luaPool  list.List
	mutex    sync.Mutex
}

type luaState struct {
	lState *glua.LState
	lModule glua.LValue
}

func IntToLuaNumber(number int64)glua.LNumber{
	return glua.LNumber(float64(number))
}

func FloatToLuaNumber(number float64)glua.LNumber{
	return glua.LNumber(number)
}

func ToLuaString(str string)glua.LString{
	return glua.LString(str)
}

func ToLuaBoolen(b bool)glua.LBool{
	return glua.LBool(b)
}

func (lua *Lua)newLuaState()(ls *luaState){
	ls = &luaState{}
	ls.lState = glua.NewState()
	ls.lState.Push(ls.lState.NewFunctionFromProto(lua.luaProto))
	if nil ==ls.lState.PCall(0, 1, nil){
		ls.lModule = ls.lState.Get(-1)
	}else{
		ls = nil
	}
	return ls
}

func (lua *Lua) lock(){
	lua.mutex.Lock()
}

func (lua *Lua) unlock(){
	lua.mutex.Unlock()
}

func (lua *Lua)getLuaState() *luaState{
	lua.lock()
	defer lua.unlock()
	var ls *luaState = nil
	if lua.luaPool.Len() <= 0{
		ls = lua.newLuaState()
		if nil == ls{
			out("getLuaState ==> nil == newLuaState")
			return nil
		}
	}else{
		e := lua.luaPool.Front()
		lua.luaPool.Remove(e)
		ls = e.Value.(*luaState)
	}
	return ls
}

func (lua *Lua)putState(ls *luaState){
	lua.lock()
	defer lua.unlock()

	lua.luaPool.PushBack(ls)
}

func (lua *Lua)Call(strFunctionName string, functionReturnCount int, functionParam ...glua.LValue)(res []interface{}, err error){
	ls := lua.getLuaState()
	if ls == nil{
		err = errors.New("getLuaState == nil")
	}else{
		for{
			luaErr := ls.lState.CallByParam(glua.P{
				Fn: ls.lState.GetField(ls.lModule, strFunctionName),
				NRet:functionReturnCount,
				Protect:true,
			}, functionParam...)

			if luaErr != nil{
				err = luaErr
				out("%s", luaErr.Error())
				debug.PrintStack()
				break
			}

			res := make([]interface{}, functionReturnCount)

			for i := functionReturnCount - 1; i >= 0; i--{
				resValue := ls.lState.Get(-1)
				switch resValue.Type() {
				case glua.LTNumber:
					res[i] = glua.LVAsNumber(resValue.(glua.LNumber))
				case glua.LTString:
					res[i] = glua.LVAsString(resValue.(glua.LString))
				case glua.LTBool:
					res[i] = glua.LVAsBool(resValue.(glua.LBool))
				default:
					res[i] = nil
				}
				ls.lState.Pop(1)
			}
			lua.putState(ls)
			break
		}
	}
	return res, err
}

func New(scriptPath string)(lua *Lua, err error){
	file, errFile := os.Open(scriptPath)
	defer file.Close()
	lua = nil
	if errFile != nil{
		err = errFile
		out("os.Open(), %s fail, %s", scriptPath, err.Error())
		return lua, err
	}else{
		reader := bufio.NewReader(file)
		chunk, _ := parse.Parse(reader, scriptPath)
		proto, err := glua.Compile(chunk, scriptPath)
		if err != nil{
			out("os.Open(), %s fail, %s", scriptPath, err.Error())
			panic(err)
		}
		lua = &Lua{}
		lua.luaProto = proto
		return lua, nil
	}
}