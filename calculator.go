package main

import (
	"fmt"
	"strconv"
)

var (
	c int
	d string
)

func add(a, b int) int {
	c = a + b
	return c
}

func subtract(a, b int) int {
	//Insert code here
	c = a - b
	return c

}

func multiply(a, b int) int {
	//Insert code here
	c = a * b
	return c

}

func divide(a, b int) int {
	//Insert code here
	//consider for b = 0
	if b == 0 {
		c = 0

	} else {
		c = a / b

	}
	return c

}

func main() {
	var a, b int
	var process string
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&a)
	fmt.Println("Enter process: (add, sub, mul, div)")
	fmt.Scanln(&process)
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&b)

	//Insert code here
	if process == "add" {
		d = strconv.Itoa(add(a, b))
	} else if process == "sub" {
		d = strconv.Itoa(subtract(a, b))
	} else if process == "div" {
		d = strconv.Itoa(divide(a, b))
	} else if process == "mul" {
		d = strconv.Itoa(multiply(a, b))
	} else {
		d = "you have keyed in an incorrect process"
	}
	fmt.Println(d)
}
