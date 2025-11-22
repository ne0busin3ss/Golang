package main

import (
	"os"
	"testing"
	"time"
)

func TestReadWord_Success(t *testing.T) {
	// Create a pipe to mock stdin
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	// Save original stdin and restore it after the test
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = r

	// Write to the pipe in a goroutine
	go func() {
		w.Write([]byte("hello\n"))
		w.Close()
	}()

	ch := make(chan result)
	go readword(ch)

	select {
	case res := <-ch:
		if res.err != nil {
			t.Errorf("Expected no error, got %v", res.err)
		}
		if res.val != "hello" {
			t.Errorf("Expected 'hello', got '%s'", res.val)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("Test timed out")
	}
}

func TestReadWord_Timeout(t *testing.T) {
	// Mock stdin but don't write anything to simulate a delay/timeout
	r, _, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = r

	ch := make(chan result)
	go readword(ch)

	select {
	case <-ch:
		t.Fatal("Expected timeout, but got a result")
	case <-time.After(100 * time.Millisecond):
		// Success: we expected a timeout (or rather, no immediate result)
		// In a real integration test with main(), we'd check context cancellation.
		// Here we just verify readword blocks as expected.
	}
}
