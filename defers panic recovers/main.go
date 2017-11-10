package main

import (
	"fmt"
)

func main() {
	//defer fmt.Println("one")
	//defer fmt.Println("two")
	fmt.Println("hello")
	testpanics()
	fmt.Println("World")
}

func testpanics() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("we recovered from a panic!")
		}
	}()

	panic("A panic happened")
}
