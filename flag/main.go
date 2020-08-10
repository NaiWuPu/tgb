package main

import (
	"flag"
	"fmt"
)

func main() {
	id := flag.Int("id", 0, "id")
	DeviceId := flag.String("p", "", "p")
	flag.Parse()
	fmt.Println(*id)
	fmt.Println(*DeviceId)
}
