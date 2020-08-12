package main

import "fmt"

const (
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBlue
)

func main() {
	fmt.Printf("%d %d %d %d\n", FlagNone, FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b %b\n", FlagNone, FlagRed, FlagGreen, FlagBlue)
}
