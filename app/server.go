package app

import (
	"github.com/toonsevrin/simplechain/types"
	"net/http"
	"github.com/gorilla/mux"

	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"
	"os"
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
			json.NewEncoder(writer).Encode(Success{false, str("Unauthorized")})
		}
	})
	router.Methods("POST").Path("/mineBlock").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if isLocalhostOrPeer(*server, *request) {
			body, err := ioutil.ReadAll(request.Body)
			if err != nil {
				json.NewEncoder(writer).Encode(Success{false, str("An error occurred reading request body")})
				fmt.Println(err.Error())
				return
			}
			data := string(body)
			block := server.App.createAndAddNextBlock(data)
			json.NewEncoder(writer).Encode(Success{Success:true})
			server.App.broadcast(block)
		}else {
			json.NewEncoder(writer).Encode(Success{false, str("Unauthorized")})
		}
	})
	router.Methods("POST").Path("/addBlock").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if isLocalhostOrPeer(*server, *request) {
			block := types.Block{}
			body, err := ioutil.ReadAll(request.Body)
			if err != nil {
				json.NewEncoder(writer).Encode(Success{false, str("An error occurred reading request body")})
				fmt.Println(err.Error())
				return
			}
			if err := json.Unmarshal(body, &block); err != nil {
				json.NewEncoder(writer).Encode(Success{false, str("An error occurred parsing request body")})
				fmt.Println(err.Error())
				return
			}
			if server.App.HasBlock(block) {
				json.NewEncoder(writer).Encode(Success{false, str("Block already exists")})
				fmt.Println("Received block that already exists in db.")
				return
			}
			if !block.IsValid() {
				json.NewEncoder(writer).Encode(Success{false, str("Received invalid block")})
				fmt.Println("Received invalid block")
				return
			}
			if uint32(len(server.App.Blockchain)) == block.Index { //next block
				if block.PreviousHash == server.App.getLatestBlock().Hash { //next block references your chain
					server.App.AddBlock(block)
					json.NewEncoder(writer).Encode(Success{Success: true})
					server.App.broadcast(block)
				} else {
					json.NewEncoder(writer).Encode(Success{Success: false, Error: str("Invalid hash")})
				}
			} else if uint32(len(server.App.Blockchain)) < block.Index { //block is in the future
				if url , exists := server.App.PeerAddresses[strings.Split(request.RemoteAddr, ":")[0]]; exists {
				RemoteChain := []types.Block{}
				response, err := http.NewRequest("GET", url +"/blocks", nil)
				if err != nil {
					json.NewEncoder(writer).Encode(Success{Success: false, Error: str(err.Error())})
					fmt.Println(err.Error())
					return
				}
				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					json.NewEncoder(writer).Encode(Success{Success: false, Error: str(err.Error())})
					fmt.Println(err.Error())
					return
				}
				if err := json.Unmarshal(body, &RemoteChain); err != nil {
					json.NewEncoder(writer).Encode(Success{Success: false, Error: str(err.Error())})
					fmt.Println(err.Error())
					return
				}
				if (server.App.pickLongestChain(RemoteChain)) {
					json.NewEncoder(writer).Encode(Success{Success: true})
					server.App.broadcast(block)
				} else {
					json.NewEncoder(writer).Encode(Success{Success: false, Error: str("Peer has a longer chain")})
				}
			}else {
					json.NewEncoder(writer).Encode(Success{Success: false, Error: str("Peer is not mutually tethered")})
				}
		}
		}else {
			json.NewEncoder(writer).Encode(Success{Success: false, Error: str("Unauthorized")})
		}
	})
	router.Methods("POST").Path("/addPeer").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if isLocalhost(*request){
			body, err := ioutil.ReadAll(request.Body)
			if err != nil {
				json.NewEncoder(writer).Encode(Success{Success: false, Error: str(err.Error())})
				fmt.Println(err.Error())
				return
			}
			peer := Peer{}
			if err := json.Unmarshal(body, &peer); err != nil {
				json.NewEncoder(writer).Encode(Success{Success: false, Error: str(err.Error())})
				fmt.Println(err.Error())
				return
			}
			server.App.Peers[peer.getUrl()] = &peer
			server.App.PeerAddresses[peer.Ip] = peer.getUrl()
			json.NewEncoder(writer).Encode(Success{Success: true})
		}else{
			writer.Write([]byte("Only localhost can add peers"))
		}
	})
	http.ListenAndServe(":" + getPort(), router)
}
func getPort() string{
	port := os.Getenv("PORT")
	if(port == ""){
		return "8080"
	}else {
		return port
	}
}
func isLocalhostOrPeer(server Server, request http.Request) bool{
	_, isPeer := server.App.PeerAddresses[request.RemoteAddr]
	return isPeer || isLocalhost(request)
}
func isLocalhost(req http.Request) bool{
	fmt.Println(req.RemoteAddr)
	return strings.Contains(req.RemoteAddr, "127.0.0.1") || strings.Contains(req.RemoteAddr, "[::1]")//::1 is ipv6
}

type Success struct {
	Success bool `json:"success"`
	Error *string `json:"error"`
}

func str(str string) *string{
	return &str
}