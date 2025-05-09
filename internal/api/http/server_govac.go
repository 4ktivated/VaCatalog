package http

import (
	"log"
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

func (s GoVacServer) ListenAndServe(addr string) {
	httpMux := http.NewServeMux()

	httpMux.HandleFunc("/sl", usecase.SuckLie)
	httpMux.HandleFunc("/lang", usecase.ApiLang)

	log.Fatal(http.ListenAndServe(addr, httpMux))
}
