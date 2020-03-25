package main

import "fmt"

var BuildTime = ""
var Debug = ""

func A(args ...interface{}){
	for k, v := range args{
		fmt.Printf("k:[%v], v:[%v]", k, v)
	}
}

func main(){
	//A(1, 2, 3)
	face := make([]interface{}, 0)
	face = append(face, 1, 2)
	A(face...)

}
