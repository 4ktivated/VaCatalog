package http

import (
	"log"
	"net/http"

	"some_app/internal/handler"
)

func Run(addr string) {

	//start sched

	httpMux := http.NewServeMux()
	warped := NewLogger(httpMux)

	httpMux.HandleFunc("/api/", handler.ApiLang)

	log.Printf("Serv start on http:://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, warped))
}
