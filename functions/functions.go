package functions

import (
	mem "github.com/ChrisCodeX/Cache-System-Go/memory"
)

func NewCache(f mem.Function) *mem.Memory {
	return &mem.Memory{
		F:     f,
		Cache: make(map[int]mem.FunctionResult),
	}
}

func GetFibonacci(cache *mem.Memory, n int) (interface{}, error) {
	value, err := cache.Get(n)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func Fibonacci(cache *mem.Memory, n int) int {
	if n <= 1 {
		return n
	}
	fb1, _ := GetFibonacci(cache, n-1)
	fb2, _ := GetFibonacci(cache, n-2)
	return fb1.(int) + fb2.(int)
}
