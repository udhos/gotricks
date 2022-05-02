package main

/*
goroutines spawn send
---------- ----- ----------
1M          2.6s      512ms
2M          4.7s    1.064s
3M         30.6s   47.781s

$ go run ./tricks/channel01 1000000
2022/05/02 00:26:54 GOMAXPROCS = 8
2022/05/02 00:26:54 count = 1000000
2022/05/02 00:26:54 spawning 1000000 goroutines...
2022/05/02 00:26:57 spawning 1000000 goroutines... elapsed: 2.617056393s
2022/05/02 00:26:57 sending value across 1000000 goroutines...
2022/05/02 00:26:58 sending value across 1000000 goroutines... elapsed: 512.41417ms

$ go run ./tricks/channel01 2000000
2022/05/02 00:27:02 GOMAXPROCS = 8
2022/05/02 00:27:02 count = 2000000
2022/05/02 00:27:02 spawning 2000000 goroutines...
2022/05/02 00:27:07 spawning 2000000 goroutines... elapsed: 4.736177983s
2022/05/02 00:27:07 sending value across 2000000 goroutines...
2022/05/02 00:27:08 sending value across 2000000 goroutines... elapsed: 1.064864313s

$ go run ./tricks/channel01 3000000
2022/05/02 00:27:15 GOMAXPROCS = 8
2022/05/02 00:27:15 count = 3000000
2022/05/02 00:27:15 spawning 3000000 goroutines...
2022/05/02 00:27:46 spawning 3000000 goroutines... elapsed: 30.665244639s
2022/05/02 00:27:46 sending value across 3000000 goroutines...
2022/05/02 00:28:34 sending value across 3000000 goroutines... elapsed: 47.781887693s
*/

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: channel01 count")
	}

	count, errConv := strconv.ParseInt(os.Args[1], 10, 64)
	if errConv != nil {
		log.Fatalf("count error: %v", errConv)
	}

	log.Printf("GOMAXPROCS = %d", runtime.GOMAXPROCS(0))
	log.Printf("count = %d", count)

	first := make(chan int)
	in := first
	var out chan int

	spawn := time.Now()
	log.Printf("spawning %d goroutines...", count)

	for i := int64(0); i < count; i++ {
		out = make(chan int)
		go forward(in, out)
		in = out
	}

	log.Printf("spawning %d goroutines... elapsed: %v", count, time.Since(spawn))

	log.Printf("sending value across %d goroutines...", count)

	begin := time.Now()
	first <- 5
	<-out

	log.Printf("sending value across %d goroutines... elapsed: %v", count, time.Since(begin))
}

func forward(in <-chan int, out chan<- int) {
	value := <-in
	out <- value
}
