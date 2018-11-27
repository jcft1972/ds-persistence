package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"crypto/md5"
	"io"
	"encoding/hex"
	)
type Document struct {
	ID string
	Name string
	Size int64
}

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

func getDocuments(w http.ResponseWriter, r *http.Request) {
	var files []Document
	var tmp Document
	

	//root := "\\gitexercises\\go.exercises\\ToursOfGo"
	root := "."
	

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		
		//id, err1 := hash_file_md5(path)
		//if err1 != nil {
		//	panic(err1)
		//}

		id := genIDmd5(path)

		tmp = Document{ID : id, Name: info.Name(), Size: info.Size()}
		files = append(files, tmp)

		return nil
	}) 

	if err != nil {
		panic(err)
	} 

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
	

}

func main () {
	router := mux.NewRouter()
	router.HandleFunc("/documents", getDocuments).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", router))
}