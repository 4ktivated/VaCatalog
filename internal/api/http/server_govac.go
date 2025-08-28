package http

import (
	"context"
	"net/http"
	"some_app/internal/metrics"
	"some_app/internal/repository"
	"some_app/internal/usecase"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type GoVacServer struct {
	logger    *zap.SugaredLogger
	pgRepo    *repository.PgRepository
	redisRepo *repository.RedisRepository
	metrics   *metrics.Metrics
}

func NewGoVacServer(ctx context.Context, logger *zap.SugaredLogger, db *repository.PgRepository, rdb *repository.RedisRepository, metrics *metrics.Metrics) GoVacServer {
	return GoVacServer{
		logger:    logger,
		pgRepo:    db,
		redisRepo: rdb,
		metrics:   metrics,
	}
}

func (s GoVacServer) ListenAndServe(ctx context.Context, addr string) error {
	httpMux := http.NewServeMux()

	httpMux.HandleFunc("/sl", usecase.SuckLie) // test
	httpMux.HandleFunc("/lang", usecase.ApiLang)

	//metrics
	httpMux.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(addr, httpMux)

}
