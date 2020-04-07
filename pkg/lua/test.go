package lua

import (
	"fmt"
	"github.com/fatezhou/go-lang"
)

func TestLua(){
	luaObject, err := zoyee_go_lang.New("../scripts/test.lua")
	fmt.Print(err)
	luaObject.Call("testLua", 0)
	res, err := luaObject.Call("add", 1, zoyee_go_lang.FloatToLuaNumber(1), zoyee_go_lang.IntToLuaNumber(2))
	fmt.Print(res, err)
}
