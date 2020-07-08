package main

import (
	"fmt"
	"time"
)

type Pp struct {
	Na int
	Ma int
}

func (P Pp)P1()  {
	P.Ma ++
	P.Na ++
	fmt.Println(P)
}

func (P *Pp) P2() {
	P.Ma ++
	P.Na ++
}


func main() {
	var P = Pp{
		Na: 1,
		Ma: 10,
	}
	P.P1()
	fmt.Println(P)
	P.P2()
	fmt.Println(P)

	fmt.Println(time.Now().UnixNano())
}
