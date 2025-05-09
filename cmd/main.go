package main

import (
	"fmt"
	"os"
	govacserver "some_app/internal/api/http"
	"sync"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load()
	printErrorAndExit(err)

	logger, err := initLogger()
	printErrorAndExit(err)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		addr := os.Getenv("APP_ADDR")
		if addr == "" {
			addr = "0.0.0.0"
		}

		port := os.Getenv("APP_PORT")
		if port == "" {
			port = "8000"
		}

		fmt.Printf("Starting server on %s:%s\n", addr, port)
		server := govacserver.NewGoVacServer(logger)
		server.ListenAndServe(fmt.Sprintf("%s:%s", addr, port))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("sync vac if app start")
	}()

	wg.Wait()
}

func initLogger() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("cant create logger error :%w", err)
	}
	return logger.Sugar(), nil

}

func printErrorAndExit(err error) {
	if err != nil {
		fmt.Printf("Error init: %s.\nFor help use -h\n", err)
		os.Exit(1)
	}
}
