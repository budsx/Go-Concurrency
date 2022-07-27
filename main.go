package main

import (
	"fmt"
	"time"
  "sync"
)

func producer(stream Stream, ch chan *Tweet,  wg *sync.WaitGroup) {
	defer close(ch)
  defer wg.Done()
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return
		}

		ch <- tweet
	}
}

func consumer(ch chan *Tweet, wg *sync.WaitGroup) {
  defer wg.Done()
	for t := range ch {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

  var wg sync.WaitGroup

	ch := make(chan *Tweet)

  wg.Add(2)
	go producer(stream, ch, &wg)

	go consumer(ch, &wg)

  wg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
