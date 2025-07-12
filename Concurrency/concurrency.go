package main

import (
	"context"
	"fmt"
	"time"
)

func readword(ch chan<- string) {
	fmt.Println("Type a word, then hit Enter.")
	var word string
	// In a real-world application, you should handle the error returned by Scanf.
	fmt.Scanf("%s", &word)
	ch <- word
}

func main() {
	// Create a context that will be cancelled after 5 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // It's important to call cancel to release resources.

	// Use a buffered channel to prevent the readword goroutine from leaking.
	ch := make(chan string, 1)
	go readword(ch)

	select {
	case word := <-ch:
		fmt.Println("Received", word)
	case <-ctx.Done():
		fmt.Println("Timeout.")
	}
}
