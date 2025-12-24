package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("From task")
	}()

	time.Sleep(time.Second)
}
