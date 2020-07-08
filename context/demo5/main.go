package main

import (
	"context"
	"fmt"
	"time"
)

// context.WithDeadline

func main() {
	d := time.Now().Add(50 * time.Millisecond)

	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("123")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
