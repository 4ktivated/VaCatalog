package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	govacserver "some_app/internal/api/http"
	"some_app/internal/metrics"
	"some_app/internal/repository"
	shed "some_app/internal/scheduler"
	"some_app/pkg/parser"
	"sync"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var signalToName = map[os.Signal]string{
	syscall.SIGINT:  "SIGINT",
	syscall.SIGTERM: "SIGTERM",
	syscall.SIGQUIT: "SIGQUIT",
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	err := godotenv.Load("cmd/.env")
	if err != nil {
		cancelFunc()
	}

	logger, err := initLogger()
	if err != nil {
		cancelFunc()
	}

	db := initPgDb(logger)
	rdb := initRedis(ctx, logger)

	osSigCh := make(chan os.Signal, 1)
	signal.Notify(osSigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for {
			s := <-osSigCh
			logger.Info("Received exit signal! Initialize graceful shutdown", zap.String("name", signalToName[s]))
			cancelFunc()
		}
	}()

	defer func() {
		if msg := recover(); msg != nil {
			err := fmt.Errorf("%s", msg)
			logger.Error("recovered from panic, but application will be terminated", zap.Error(err))
			cancelFunc()
		}
	}()

	pgRepo := repository.NewRepository(db, logger)
	redisRepo := repository.NewRedisRepo(rdb, logger)
	metric := metrics.New()

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
		server := govacserver.NewGoVacServer(ctx, logger, pgRepo, redisRepo, metric)
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

	wg.Add(1)
	go func() {
		<-ctx.Done()
		rdb.Close()
		db.Close()
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

func initPgDb(logger *zap.SugaredLogger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Fatal("can't init redis client", zap.Error(err))
	}
	return dbpool

}

func initRedis(ctx context.Context, logger *zap.SugaredLogger) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logger.Fatal("can't init redis client", zap.Error(err))
	}

	return rdb
}
