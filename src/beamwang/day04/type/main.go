package main

import "fmt"

type myInt int

func main() {
	var n myInt = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)
}
