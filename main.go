package main

import (
	"fmt"
	"log"
	"time"
)

func Fibonacci(cache *Memory, n int) int {
	if n <= 1 {
		return n
	}
	fb1, _ := GetFibonacci(cache, n-1)
	fb2, _ := GetFibonacci(cache, n-2)
	return fb1.(int) + fb2.(int)
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
}

type Function func(cache *Memory, key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func GetFibonacci(cache *Memory, n int) (interface{}, error) {
	value, err := cache.Get(n)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	result, exists := m.cache[key]
	if !exists {
		result.value, result.err = m.f(key)
		m.cache[key] = result
	}
	return result.value, result.err
}

func main() {
	startRun := time.Now()
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38}
	for _, n := range fibo {
		start := time.Now()
		value, err := cache.Get(n)
		fmt.Printf("%d,%s,%d\n", n, time.Since(start), value)
		if err != nil {
			log.Println(err)
		}
	}
	fmt.Printf("Process completed in %s\n", time.Since(startRun))
}
