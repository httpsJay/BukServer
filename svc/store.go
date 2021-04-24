package svc

import (
	"fmt"
	"regexp"
)

type SUFFIXChanMessage struct {
	Key        string
	ReturnChan chan []string
}

type PREFIXChanMessage struct {
	Key        string
	ReturnChan chan []string
}

type GETChanMessage struct {
	Key        string
	ReturnChan chan string
}

type DataBase struct {
	Data       map[string]string
	SETChan    chan map[string]string
	GETChan    chan GETChanMessage
	PREFIXChan chan PREFIXChanMessage
	SUFFIXChan chan SUFFIXChanMessage
}

// chanListener will listen to all request for DB access
// As golang is not thread-safe, It will prevent issues by using channels
func (db *DataBase) chanListener() {
	for {
		select {

		//set <key><value> related channel
		case item := <-db.SETChan:
			key := item["key"]
			value := item["value"]
			db.Data[key] = value
			fmt.Println(db.Data)

		//get <key> related channel
		case item := <-db.GETChan:
			key := item.Key
			returnChan := item.ReturnChan
			if value, ok := db.Data[key]; ok {
				returnChan <- value
			} else {
				returnChan <- "ERROR: NOT FOUND"
			}

		//search with prefix related channel
		case item := <-db.PREFIXChan:
			filterStr := item.Key
			returnChan := item.ReturnChan

			res := new(ListResponse)
			for key := range db.Data {
				if matched, _ := regexp.MatchString(filterStr+".*", key); matched {
					res.List = append(res.List, key)
				}
			}
			returnChan <- res.List

		//search with suffix related channel
		case item := <-db.SUFFIXChan:
			filterStr := item.Key
			returnChan := item.ReturnChan

			res := new(ListResponse)
			for key := range db.Data {
				if matched, _ := regexp.MatchString(".*"+filterStr, key); matched {
					res.List = append(res.List, key)
				}
			}
			returnChan <- res.List

		}
	}
}

// Create DataBase instance
func Create() *DataBase {
	db := &DataBase{
		Data:       make(map[string]string),           // Key:value store
		SETChan:    make(chan map[string]string, 250), // Buffered-channel
		GETChan:    make(chan GETChanMessage),         // UnBuffered-channel
		PREFIXChan: make(chan PREFIXChanMessage),      // UnBuffered-channel
		SUFFIXChan: make(chan SUFFIXChanMessage),      // UnBuffered-channel
	}
	go db.chanListener()
	return db
}
