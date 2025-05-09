package scheduler

import (
	"fmt"
	"some_app/pkg/parser"
	"time"
)

type Scheduler interface {
	InitSync()
	RunSync()
	sync()
}

type ShedulerPars struct {
	parseClient *parser.ParseClient
}

func NewShedilerPars(parseClient *parser.ParseClient) *ShedulerPars {
	return &ShedulerPars{
		parseClient: parseClient,
	}
}

// TODO: подумать над планировщиком
func (s *ShedulerPars) RunSync() {
	client := parser.NewParseClient()

	ticker := time.NewTicker(20 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				client.HHparser()
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func (s *ShedulerPars) InitSync() error {
	//do somthing
	return nil
}

func sync() {
	// do somthing
}
