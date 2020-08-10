package main

func main() {
	var c = make(chan int64, 20)

	for {
		<-c
	}
}
