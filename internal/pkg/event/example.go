package event

import "fmt"

func test()  {
	Event.On("Hello", func(args ...interface{}) {
		fmt.Println("Hello")
		for _, v := range args  {
			fmt.Printf("%+v\n", v)
		}
	})

	Event.On("Bye", func(args ...interface{}) {
		fmt.Println("Bye")
		for _, v := range args  {
			fmt.Printf("%+v\n", v)
		}
	})

	Emit("Hello", "ABC", 1234)
	Emit("Bye", "88", 88)
}
