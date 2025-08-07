package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Service 1: simple binary calculator")
	fmt.Println("Usage: <op> x y")
	fmt.Println("where <op> is one of 'add', 'sub', 'mul', 'div'; x and y are integers")

	op := os.Args[1]

	x, _ := strconv.Atoi(os.Args[2])
	y, _ := strconv.Atoi(os.Args[3])

	var ans int 

	switch op {
	case "add": 
		ans = add(x, y)
	case "sub": 
		ans = sub(x, y)
	case "mul":
		ans = mul(x, y)
	case "div": 
		ans = div(x, y)
	default: 
		fmt.Println("What the hell :)")
	}

	fmt.Printf("= %d", ans)
}

func add(x int, y int) int {
	return x + y 
}

func sub(x int, y int) int {
	return x - y
}

func mul(x int, y int) int { 
	return x * y
}

func div(x int, y int) int {
	// not handle divide by 0 
	return x/y 
}