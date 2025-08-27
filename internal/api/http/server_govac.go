package http

import (
	"context"
	"net/http"
	"some_app/internal/usecase"

	"go.uber.org/zap"
)

type GoVacServer struct {
	logger *zap.SugaredLogger
}

func NewGoVacServer(logger *zap.SugaredLogger) GoVacServer {
	return GoVacServer{logger: logger}
}

func (s GoVacServer) ListenAndServe(ctx context.Context, addr string) error {
	httpMux := http.NewServeMux()

	httpMux.HandleFunc("/sl", usecase.SuckLie) // test
	httpMux.HandleFunc("/lang", usecase.ApiLang)

	return http.ListenAndServe(addr, httpMux)

}
