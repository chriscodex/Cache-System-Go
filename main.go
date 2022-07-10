package main

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f Function
}

type Function func(key int) (interface{}, error)

func main() {
	f := Fibonacci(5)
	fmt.Println(f)
}
