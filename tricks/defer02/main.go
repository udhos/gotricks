package main

import (
	"log"
	"time"
)

func main() {
	bigWork()
}

func bigWork() {
	// the trace() function is executed immediately
	// but the function returned from trace is executed when this function returns
	defer trace("bigWork")()
	log.Printf("bigWork: doing a lot of work")
	time.Sleep(3 * time.Second)
}

// used with defer trace("")() to trace func execution
func trace(msg string) func() {
	log.Printf("trace begin: %s", msg) // executed immediately at defer call
	start := time.Now()
	return func() {
		// executed when the caller function returns
		log.Printf("trace end: %s - elapsed: %v", msg, time.Since(start))
	}
}
