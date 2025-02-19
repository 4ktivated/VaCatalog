package http

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
    handler http.Handler
}


func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    l.handler.ServeHTTP(w, r)
    log.Printf("%s %s %s %f sec ", r.Method, r.RemoteAddr, r.URL.Path, time.Since(start).Seconds())
}

func NewLogger(handlerToWrap http.Handler) *Logger {
    return &Logger{handlerToWrap}
}