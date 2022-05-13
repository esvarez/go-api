package api

import (
	"github.com/esvarez/go-api/api/handler"
	"github.com/gorilla/mux"
	"log"
)

const (
	Port = "4200"
)

func Start() {
	var (
		router    = mux.NewRouter()
		bgHandler = handler.NewBoardGameHandler(nil)
		server    = newServer(Port, router)
	)

	handler.MakeBoardGameHandler(router, bgHandler)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("%v error starting server", err)
	}
}
