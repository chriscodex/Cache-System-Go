package main

import (
	"fmt"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
}

type Function func(key int) interface{}

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) interface{} {
	result, exist := m.cache[key]
	if !exist {
		result.value = m.f(key)
		m.cache[key] = result
	}
	return result.value
}

func GetFibonacci(n int) interface{} {
	return Fibonacci(n)
}

func main() {
	//startRun := time.Now()
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38}
	for _, n := range fibo {
		start := time.Now()
		value := cache.Get(n)
		fmt.Printf("%d,%s,%d\n", n, time.Since(start), value)
	}
	fmt.Println("Process Completed")
}
