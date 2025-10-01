package main

import (
	"fmt"
)

const year int = 2025

func main() {
	var name string = "hari"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name) //place holder %T returns the datatype

	var isloggedin bool = false
	fmt.Println(isloggedin)
	fmt.Printf("The type is: %T \n", isloggedin)

	var age uint8 = 22 // 0-255 unsigned integer
	fmt.Println(age)
	fmt.Printf("The type is %T \n", age)

	name1 := "Hari Prasath" // := walrus operator could be only used locally not globally
	fmt.Println(name1)

	fmt.Println(year)

	const company = "CEVA"
	fmt.Println(company)
}
