package main

import "fmt"

func process(x any) {
	switch v := x.(type) {
	case int :
		fmt.Println("int", v)
		
	case bool :
		fmt.Println("bool", v)

	case string :
		fmt.Println("string", v)

	default:
		fmt.Println("bad value")
	}
	
}

func main() {
	process(10)
	process(10.1)
	process("hello")
	process(true)
}