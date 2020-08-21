package main

import (
	"fmt"
	"time"
)

func main() {
	sec := time.NewTimer(0)
	defer sec.Stop()
	for {
		select {
		case i := <-sec.C:
			fmt.Println(i)
			sec.Reset(20 * time.Second)
		}
	}

}
