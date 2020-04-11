package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Finished bool   `json:"finished,omitempty"`
}

var taskList []task

func getTaskEndPoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusFound)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskList)
}

func main() {
	router := mux.NewRouter()
	taskList = []task{}

	router.HandleFunc("/task", getTaskEndPoint)

	log.Fatal(http.ListenAndServe(":3000", router))
}
