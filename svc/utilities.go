package svc

import (
	"bukukas/store"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

// Create a global package level In-Mem DB map
var db = store.Create()

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
	fmt.Println("searchWithSuffix()")
	vars := mux.Vars(r)
	filterStr := vars["str"]

	fmt.Println("suffix : ", filterStr)

	res := new(ListResponse)
	for key := range db.Data {
		if matched, _ := regexp.MatchString(".*"+filterStr, key); matched {
			res.List = append(res.List, key)
		}
	}
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
	fmt.Println("searchWithPrefix()")
	vars := mux.Vars(r)
	filterStr := vars["str"]

	fmt.Println("prefix : ", filterStr)

	res := new(ListResponse)
	for key := range db.Data {
		if matched, _ := regexp.MatchString(filterStr+".*", key); matched {
			res.List = append(res.List, key)
		}
	}
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

	value := db.Get(key)

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

// Used for /set route
func setKeyValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]

	fmt.Println("\nSet Key : ", key, ", Value : ", value)

	db.Set(key, value)
	fmt.Printf("DataBase Data : %#v", db.Data)
	return
}
