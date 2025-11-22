package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func readword(ch chan<- string) {
	fmt.Println("Type a word, then hit Enter.")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		ch <- scanner.Text()
	} else {
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
		}
		// If EOF or error, we might want to close or send a signal,
		// but for this simple example, we'll just return.
	}
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
