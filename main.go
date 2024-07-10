package main

import (
	"context"
	"log"
	"some_code/app"
)

func main() {
	err := app.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
