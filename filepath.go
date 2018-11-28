package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"net/http"
	"crypto/md5"
	"io"
	"encoding/hex"
	"github.com/gorilla/mux"
	)
type Document struct {
	ID string
	Name string
	Size int64
}

//root := "\\gitexercises\\go.exercises\\ToursOfGo"
var root = "."

func hash_file_md5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

func genIDmd5(fileOpen string) string{
	hasher := md5.New()
	hasher.Write([]byte(fileOpen))
	return string(hex.EncodeToString(hasher.Sum(nil)))
}


func listFiles(root string) ([]Document, error){

	var files []Document
	var tmp Document

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		id := genIDmd5(path)
		tmp = Document{ID : id, Name: info.Name(), Size: info.Size()}
		files = append(files, tmp)
		return nil
	}) 
	return files, err

}

func getDocuments(w http.ResponseWriter, r *http.Request) {
	
	files, err := listFiles(root)

	if err != nil {
		panic(err)
	} 

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
	
}

func GetDocumentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	files, err := listFiles(root)
	
	if err != nil {
		panic(err)
	} 
	
	for _, file := range files {
		if file.ID == params["id"] {
			json.NewEncoder(w).Encode(file)
			return
		}
	}

	json.NewEncoder(w).Encode(&Document{})
}

func CreateDocumentById(w http.ResponseWriter, r *http.Request){}
func DeletetDocumentById(w http.ResponseWriter, r *http.Request){}


// func main () {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/documents", getDocuments).Methods("GET")
// 	log.Fatal(http.ListenAndServe(":9000", router))
// }