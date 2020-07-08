package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()
	nextYear, err := time.Parse("2006-01-02","2020-06-23")
	log.Println(err)
	log.Println(nextYear)

	d := nextYear.Sub(now)
	fmt.Println(d)
}