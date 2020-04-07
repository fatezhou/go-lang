package event

import (
	"fmt"
	"github.com/fatezhou/go-lang"
)

func test()  {
	zoyee_go_lang.Event.On("Hello", func(args ...interface{}) {
		fmt.Println("Hello")
		for _, v := range args  {
			fmt.Printf("%+v\n", v)
		}
	})

	zoyee_go_lang.Event.On("Bye", func(args ...interface{}) {
		fmt.Println("Bye")
		for _, v := range args  {
			fmt.Printf("%+v\n", v)
		}
	})

	zoyee_go_lang.Emit("Hello", "ABC", 1234)
	zoyee_go_lang.Emit("Bye", "88", 88)
}
