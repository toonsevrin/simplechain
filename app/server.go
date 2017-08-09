package app

import (
	"github.com/toonsevrin/simplechain/types"
	"net/http"
	"github.com/gorilla/mux"

	"encoding/json"
	"io/ioutil"
	"github.com/prometheus/common/log"
	"fmt"
)

type Server struct {
	App App
}
func (server *Server) Init(){

	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/blocks").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if isLocalhostOrPeer(*server, *request){
			json.NewEncoder(writer).Encode(server.App.Blockchain)
		}else {
			json.NewEncoder(writer).Encode(Success{false, "Unauthorized"})
		}
	})
	router.Methods("POST").Path("/mineBlock").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if isLocalhostOrPeer(*server, *request) {
			body, err := ioutil.ReadAll(request.Body)
			if err != nil {
				json.NewEncoder(writer).Encode(Success{false, "An error occurred reading request body"})
				log.Error(err.Error())
				return
			}
			data := string(body)
			block := server.App.createAndAddNextBlock(data)
			json.NewEncoder(writer).Encode(Success{Success:true})
			server.App.broadcast(block)
			log.Error("log")
		}else {
			json.NewEncoder(writer).Encode(Success{false, "Unauthorized"})
		}
	})
	router.Methods("POST").Path("/addBlock").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if isLocalhostOrPeer(*server, *request){
			block := types.Block{}
			body, err := ioutil.ReadAll(request.Body)
			if err != nil {
				json.NewEncoder(writer).Encode(Success{false,"An error occurred reading request body"})
				log.Error(err.Error())
				return
			}
			if err := json.Unmarshal(body, block); err != nil {
				json.NewEncoder(writer).Encode(Success{false,"An error occurred parsing request body"})
				log.Error(err.Error())
				return
			}
			if server.App.HasBlock(block){
				json.NewEncoder(writer).Encode(Success{false,"Block already exists"})
				log.Debug("Received block that already exists in db.")
				return
			}
			if !block.IsValid() {
				json.NewEncoder(writer).Encode(Success{false,"Received invalid block"})
				log.Debug("Received invalid block")
				return
			}
			if uint32(len(server.App.Blockchain)) == block.Index {//next block
				if block.PreviousHash == server.App.getLatestBlock().Hash {//next block references your chain
					server.App.AddBlock(block)
					json.NewEncoder(writer).Encode(Success{Success: true})
					server.App.broadcast(block)
				}
			}else if uint32(len(server.App.Blockchain)) < block.Index {//block is in the future
				RemoteChain := []types.Block{}
				response, err := http.NewRequest("GET", server.App.Peers[request.RemoteAddr].getUrl() + "/blocks", nil)
				if err != nil {
					json.NewEncoder(writer).Encode(Success{Success: false, Error: string(err.Error())})
					log.Error(err.Error())
					return
				}
				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					json.NewEncoder(writer).Encode(Success{Success: false, Error: string(err.Error())})
					log.Error(err.Error())
					return
				}
				if err := json.Unmarshal(body, RemoteChain); err != nil {
					json.NewEncoder(writer).Encode(Success{Success: false, Error: string(err.Error())})
					log.Error(err.Error())
					return
				}
				if(server.App.pickLongestChain(RemoteChain)){
					json.NewEncoder(writer).Encode(Success{Success: true})
					server.App.broadcast(block)
				}
			}
		}else {
			json.NewEncoder(writer).Encode(Success{Success: false, Error: "Unauthorized"})
		}
	})
	router.Methods("POST").Path("/addPeer").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if isLocalhost(*request){
			body, err := ioutil.ReadAll(request.Body)
			if err != nil {
				json.NewEncoder(writer).Encode(Success{Success: false, Error: string(err.Error())})
				log.Error(err.Error())
				return
			}
			peer := Peer{}
			if err := json.Unmarshal(body, peer); err != nil {
				json.NewEncoder(writer).Encode(Success{Success: false, Error: string(err.Error())});
				log.Error(err.Error())
				return
			}
			server.App.Peers[peer.getUrl()] = peer
			server.App.PeerAddresses[peer.Ip] = true
			json.NewEncoder(writer).Encode(Success{Success: true});
		}else{
			writer.Write([]byte("Only localhost can add peers"))
		}
	})
	http.ListenAndServe(":8080", router)
}
func isLocalhostOrPeer(server Server, request http.Request) bool{
	_, isPeer := server.App.PeerAddresses[request.RemoteAddr];
	return isPeer || isLocalhost(request)
}
func isLocalhost(req http.Request) bool{
	return req.RemoteAddr == "127.0.0.1"
}

type Success struct {
	Success bool `json:"success"`
	Error string `json:"error"`
}