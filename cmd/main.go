package main

import (
	test "zoyee-tool/internal/pkg/test"
)

var BuildTime = ""
var Debug = ""

func main(){
	test.Test.ParseCmd()
	test.Test.Run()
}
