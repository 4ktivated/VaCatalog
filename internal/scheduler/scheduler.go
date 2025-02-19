package scheduler

import (
	"fmt"
	"some_app/pkg/parser"
	"time"
)

// TODO: подумать над планировщиком
func Sokolas() {
	ticker := time.NewTicker(20 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				client := parser.HTTPParseClient()
				client.R.HHparser()
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
