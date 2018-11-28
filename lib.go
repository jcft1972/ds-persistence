package main

import (
	"os"
	"path/filepath"
	"crypto/md5"
	"io"
	"encoding/hex"
)


var root = "."

func check(e error) {
    if e != nil {
        panic(e)
    }
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