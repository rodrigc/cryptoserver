package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

func main() {
	port := defaultHTTPPort
	// Look up HTTP_PORT environment variable, otherwise
	// use default port specified in settings.go
	if val, ok := os.LookupEnv("HTTP_PORT"); ok {
		port = val
	}

	r := mux.NewRouter()
	r.HandleFunc("/currency/{symbol}", CurrencyHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	fmt.Printf("Starting crypto-app on %s\n", srv.Addr)
	srv.ListenAndServe()
}
