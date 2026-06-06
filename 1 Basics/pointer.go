package main

import "fmt"

func main() {
	x := 10

	ptr := &x

	fmt.Println(*ptr)
	fmt.Println(ptr)
}
