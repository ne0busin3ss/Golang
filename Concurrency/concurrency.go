package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

type result struct {
	val string
	err error
}

func readword(ch chan<- result) {
	fmt.Println("Type a word, then hit Enter.")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		ch <- result{val: scanner.Text()}
	} else {
		if err := scanner.Err(); err != nil {
			ch <- result{err: err}
		} else {
			// EOF
			ch <- result{err: fmt.Errorf("EOF")}
		}
	}
}

func main() {
	// Create a context that will be cancelled after 5 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // It's important to call cancel to release resources.

	// Use a buffered channel to prevent the readword goroutine from leaking.
	ch := make(chan result, 1)
	go readword(ch)

	select {
	case res := <-ch:
		if res.err != nil {
			fmt.Println("Error reading input:", res.err)
		} else {
			fmt.Println("Received", res.val)
		}
	case <-ctx.Done():
		fmt.Println("Timeout.")
	}
}
