package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func count() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		i := 0
		for ; i < 1e6; i++ {
		}
		_ = i
	}
}

func main() {
	fmt.Println("Version", runtime.Version())
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go count()
	}
	wg.Wait()
}
