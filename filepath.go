package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	)
type Document struct {
	ID string
	Name string
	Size int64
}

//root := "\\gitexercises\\go.exercises\\ToursOfGo"


func getDocuments(w http.ResponseWriter, r *http.Request) {
	
	files, err := listFiles(root)

	check(err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
	
}

func GetDocumentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	files, err := listFiles(root)
	
	check(err)
	
	for _, file := range files {
		if file.ID == params["id"] {
			json.NewEncoder(w).Encode(file)
			return
		}
	}

	json.NewEncoder(w).Encode(&Document{})
}


func CreateDocumentById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	file, err := cpFile(params["id"])
	check(err)
	json.NewEncoder(w).Encode(file)
}

func DeletetDocumentById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	err := rmFile(params["id"])
	check(err)
	json.NewEncoder(w).Encode(&Document{})
}


// func main () {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/documents", getDocuments).Methods("GET")
// 	log.Fatal(http.ListenAndServe(":9000", router))
// }