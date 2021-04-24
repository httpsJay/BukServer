package svc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Create a global package level In-Mem DB map
var db = Create()

// Dummy Home url
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// used for /get route
type KeyResponse struct {
	Value string
}

// used for /search route
type ListResponse struct {
	List []string
}

// Search route with suffix filter
func searchWithSuffix(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nsearchWithSuffix()")
	vars := mux.Vars(r)
	filterStr := vars["str"]

	fmt.Println("suffix : ", filterStr)

	c := make(chan []string)
	m := PREFIXChanMessage{filterStr, c}
	// sending using channel
	db.PREFIXChan <- m

	res := new(ListResponse)
	// receiving response from returnChannel
	res.List = <-c
	close(c)

	fmt.Println("response : ", res)

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// Search route with prefix filter
func searchWithPrefix(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nsearchWithPrefix()")
	vars := mux.Vars(r)
	filterStr := vars["str"]

	fmt.Println("prefix : ", filterStr)

	c := make(chan []string)
	m := PREFIXChanMessage{filterStr, c}
	// sending using channel
	db.PREFIXChan <- m

	res := new(ListResponse)
	// receiving response from returnChannel
	res.List = <-c
	close(c)

	fmt.Println("response : ", res)

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// used for /get route, returns value of the key from in-mem db
func getKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	fmt.Printf("\nget - Key [%s]", key)

	c := make(chan string)
	m := GETChanMessage{key, c}
	db.GETChan <- m
	value := <-c
	close(c)

	if value == "ERROR: NOT FOUND" {
		http.NotFound(w, r)

	} else {
		// fmt.Fprint(w, value)
		res := new(KeyResponse)
		res.Value = value

		js, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

// Used for /set route
func setKeyValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]

	fmt.Println("\nSet ", key, " -> ", value)

	m := make(map[string]string)
	m["key"] = key
	m["value"] = value
	db.SETChan <- m

	return
}
