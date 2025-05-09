package main

import (
	"context"
	"fmt"
	"os"
	govacserver "some_app/internal/api/http"
	shed "some_app/internal/scheduler"
	"some_app/pkg/parser"
	"sync"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	fmt.Println(cancelFunc) //TODO: что-то с жтимм сделать

	err := godotenv.Load()
	printErrorAndExit(err)

	logger, err := initLogger()
	printErrorAndExit(err)

	//init pool os parsers
	hhParser := parser.NewHHparser([]string{"php", "python", "golang"})

	prserClient := parser.NewParseClient(hhParser)

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
		sheduler := shed.NewShedilerPars(*logger, prserClient)
		err := sheduler.SyncOnce()
		if err != nil {
			logger.Error("cant sync vacancy for the first time")
		}
		sheduler.RunSync(ctx)
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
