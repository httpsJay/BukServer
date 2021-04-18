package svc

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Port = "8877"

func HandleRequests() {
	// Router for the server
	myRouter := mux.NewRouter().StrictSlash(true)

	// dummy home route
	myRouter.HandleFunc("/", homePage)

	// get route receives a key and return value
	myRouter.HandleFunc("/get/{key}", getKey).Methods("GET")

	// set add a dictionary entry key:value
	myRouter.HandleFunc("/set/{key}/{value}", setKeyValue).Methods("GET")

	// search can use two filter prefix and suffix. Returns regex matched list result keys
	myRouter.Path("/search").Queries("prefix", "{str}").HandlerFunc(searchWithPrefix).Methods("GET")
	myRouter.Path("/search").Queries("suffix", "{str}").HandlerFunc(searchWithSuffix).Methods("GET")

	//Starting Server using Port
	fmt.Println("localhost:8877")
	log.Fatal(http.ListenAndServe(":"+Port, myRouter))
}
