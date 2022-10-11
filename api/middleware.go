package api

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// RecoveryMiddleware recovers any panics further down the call and responds with InternalServerError
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				e := fmt.Errorf("%v", err)
				ErrorResponse(w, InternalServerError(e))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// LogResponseWriter Satisfies the http.ResponseWriter interface and saves the status code
type LogResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (lrw *LogResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// LogMiddleware logs information of the request to normal log
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		lrw := &LogResponseWriter{w, http.StatusOK}
		next.ServeHTTP(lrw, r)
		duration := time.Since(t)
		log.Println(r.Method, r.URL.Path, lrw.StatusCode, http.StatusText(lrw.StatusCode), duration)
	})
}
