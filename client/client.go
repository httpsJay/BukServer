package main

import (
	"bukukas/svc"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	host = "http://localhost"
	port = "8877"
)

var url = host + ":" + port + "/"

func Search(param, filterStr string) ([]string, error) {
	fmt.Println(url + "search?" + param + "=" + filterStr)
	r, err := http.Get(url + "search?" + param + "=" + filterStr)
	if err != nil {
		fmt.Println("Error searching request response : ", err)
		return nil, err
	}
	defer r.Body.Close()

	target := new(svc.ListResponse)
	json.NewDecoder(r.Body).Decode(&target)
	fmt.Println("received List : ", target.List)

	return target.List, nil
}

func SearchPrefix(filterStr string) ([]string, error) {
	return Search("prefix", filterStr)
}
func SearchSuffix(filterStr string) ([]string, error) {
	return Search("suffix", filterStr)
}

func Get(key string) (string, error) {
	fmt.Println(url + "get/" + key)
	r, err := http.Get(url + "get/" + key)
	if err != nil {
		fmt.Println("Error while set request response : ", err)
		return "", err
	}
	defer r.Body.Close()

	target := new(svc.KeyResponse)
	json.NewDecoder(r.Body).Decode(&target)
	fmt.Println("returnedvalue : ", target.Value)

	return target.Value, nil
}

func Set(key, value string) error {
	fmt.Println(url + "set/" + key + "/" + value)
	_, err := http.Get(url + "set/" + key + "/" + value)
	if err != nil {
		fmt.Println("Error while set request response : ", err)
		return err
	}
	return nil
}

// func main() {
// 	err := Set("jay", "roy")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	err = Set("jack", "qwer")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	err = Set("rameshk", "power")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	value, err := Get("jay")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println("Found value : ", value)

// 	resList, err := SearchPrefix("j")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println("prefix received List : ", resList)

// 	resList, err = SearchPrefix("ja")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println("prefix received List : ", resList)

// 	resList, err = SearchSuffix("k")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println("prefix received List : ", resList)

// }
