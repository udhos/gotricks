package main

/*
1M goroutines is easy:

go run ./tricks/channel01 1000000
2022/05/01 23:35:34 GOMAXPROCS = 8
2022/05/01 23:35:34 count = 1000000
2022/05/01 23:35:37 elapsed: 546.507183ms

2M goroutines is ok (time is 2x 1M):

go run ./tricks/channel01 2000000
2022/05/01 23:35:51 GOMAXPROCS = 8
2022/05/01 23:35:51 count = 2000000
2022/05/01 23:35:56 elapsed: 1.022604209s

3M goroutines is hard (time is 120x 1M):

go run ./tricks/channel01 3000000
2022/05/01 23:36:00 GOMAXPROCS = 8
2022/05/01 23:36:00 count = 3000000
2022/05/01 23:38:00 elapsed: 1m4.089638301s
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

	for i := int64(0); i < count; i++ {
		out = make(chan int)
		go forward(in, out)
		in = out
	}

	begin := time.Now()
	first <- 5
	<-out
	log.Printf("elapsed: %v", time.Since(begin))
}

func forward(in <-chan int, out chan<- int) {
	value := <-in
	out <- value
}
