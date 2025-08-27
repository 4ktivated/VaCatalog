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
	fmt.Printf("this is func: %p\n", cancelFunc) //TODO что-то сделатьс этим

	err := godotenv.Load("cmd/.env")
	if err != nil {
		cancelFunc()
	}

	logger, err := initLogger()
	if err != nil {
		cancelFunc()
	}

	//init pool os parsers
	hhParser := parser.NewHHparser([]string{"php", "python", "golang"})

	parserClient := parser.NewParseClient(hhParser)

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
		err := server.ListenAndServe(ctx, fmt.Sprintf("%s:%s", addr, port))
		if err != nil {
			cancelFunc()

		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sheduler := shed.NewShedilerPars(logger, parserClient)
		err := sheduler.SyncOnce(ctx, logger)
		if err != nil {
			logger.Error("cant sync vacancy for the first time")
		} else {
			logger.Info("all good")
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
