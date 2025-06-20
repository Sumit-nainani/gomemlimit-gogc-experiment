package server

import (
	"goenvs/handler"
	"net/http"
)

func Server() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HeavyHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return server
}
