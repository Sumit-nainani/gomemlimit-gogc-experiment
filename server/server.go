package server

import (
	"goenvs/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.HeavyHandler)
	return router
}
