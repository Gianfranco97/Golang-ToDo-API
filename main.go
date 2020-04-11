package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Finished bool   `json:"finished,omitempty"`
}

type errorMessage struct {
	Message string `json:"message,omitempty"`
}

var taskList []task

func getTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskList)
}

func getOneTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorMessage{Message: "The requested ID is not valid"})
		return
	}

	for _, item := range taskList {
		if item.ID == id {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task{})
}

func main() {
	router := mux.NewRouter()
	taskList = []task{}

	router.HandleFunc("/task", getTaskEndPoint)
	router.HandleFunc("/task/{id}", getOneTaskEndPoint)

	log.Fatal(http.ListenAndServe(":3000", router))
}
