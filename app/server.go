package app

import (
	"net/http"
	"github.com/gorilla/mux"

	"encoding/json"
)

type Server struct {
	App App
}
func (server *Server) Init(){

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/blocks").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode(server.App.Blockchain)
	})

	http.ListenAndServe(":8080", router)
}

