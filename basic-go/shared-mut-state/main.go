package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	n  int
}

func main() {
	c := &Counter{}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		c.mu.Lock()
		c.n++
		c.mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		c.mu.Lock()
		c.n++
		c.mu.Unlock()
	}()

	wg.Wait()
	fmt.Println(c.n)
}
