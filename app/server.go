package app

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Server struct {
	App App
}
func (server *Server) init(){
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/blocks").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode(server.App.Blockchain)
	})

	http.ListenAndServe(":8080", router)
}

