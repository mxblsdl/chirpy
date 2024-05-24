package main

import (
	"fmt"
	"net/http"
)

type CustomHandler struct{}

func (h *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Write the response body
	w.Write([]byte("OK"))
}

func main() {
	servMux := http.NewServeMux()

	servMux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("app"))) )
	
	customHandler := &CustomHandler{}
	servMux.Handle("/healthz", customHandler)

	serv := &http.Server{
		Addr:":8080",
		 Handler: servMux,
		}

	if err := serv.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}