package lua

import "fmt"

func TestLua(){
	luaObject, err := New("../scripts/test.lua")
	fmt.Print(err)
	luaObject.Call("testLua", 0)
	res, err := luaObject.Call("add", 1, FloatToLuaNumber(1), IntToLuaNumber(2))
	fmt.Print(res, err)
}
