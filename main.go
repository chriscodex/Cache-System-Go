package main

import (
	"fmt"
	"time"

	fn "github.com/ChrisCodeX/Cache-System-Go/functions"
	mem "github.com/ChrisCodeX/Cache-System-Go/memory"
)

func main() {
	fibo := []int{42, 40, 41, 42, 38}
	startRun := time.Now()
	cache := fn.NewCache(func(cache *mem.Memory, n int) (interface{}, error) {
		return fn.Fibonacci(cache, n), nil
	})
	for i := 0; i < len(fibo); i++ {
		n := fibo[i]
		start := time.Now()
		value, err := fn.GetFibonacci(cache, n)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("%d, %s, %d\n", n, time.Since(start), value)
	}
	fmt.Println(cache.Cache)
	fmt.Printf("Process completed in %s\n", time.Since(startRun))
}
