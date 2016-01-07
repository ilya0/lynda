package main

import "fmt"

func main() {
	var x rune = 'a' // rune is an alias for int32; normally omitted in this statement
	var y int32 = 'b'
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(string(x))
	fmt.Println(string(y))
	// conversion: rune to string
}




/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/