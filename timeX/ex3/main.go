package main

import (
	"fmt"
	"time"
)

func main() {
	sec := time.NewTimer(time.Second)
	defer sec.Stop()
	for {
		select {
		case i := <-sec.C:
			fmt.Println(i)
			sec.Reset(time.Second)
		}
	}

}
