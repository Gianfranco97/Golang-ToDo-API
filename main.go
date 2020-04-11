package main

import (
	"encoding/json"
	"io/ioutil"
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(taskList)
}

func addOneTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage{Message: "The task data is not valid"})
		return
	}

	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(taskList) + 1

	taskList = append(taskList, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func getOneTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage{Message: "The requested ID is not valid"})
		return
	}

	for _, item := range taskList {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(task{})
}

func deleteOneTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage{Message: "The requested ID is not valid"})
		return
	}

	for index, item := range taskList {
		if item.ID == id {
			taskList = append(taskList[:index], taskList[index+1:]...)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(task{})
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(task{})
}

func main() {
	router := mux.NewRouter()
	taskList = []task{}

	router.HandleFunc("/task", getTaskEndPoint).Methods("GET")
	router.HandleFunc("/task", addOneTaskEndPoint).Methods("POST")
	router.HandleFunc("/task/{id}", getOneTaskEndPoint).Methods("GET")
	router.HandleFunc("/task/{id}", deleteOneTaskEndPoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
