package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-t.C:
			fmt.Println(2)
		}
	}
}
