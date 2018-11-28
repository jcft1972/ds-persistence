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

func cpFile(path string) (Document, error) {
	var nameFile string
	fileIn, err := os.Open(path)
	check(err)
	nameFile = "Copy-" + fileIn.Name()
	defer fileIn.Close()
	fileOut, err := os.Create(nameFile)
	check(err)
	defer fileOut.Close()
	_, err = io.Copy(fileOut, fileIn)
	check(err)

	return Document{ID: genIDmd5(fileOut.Name()), Name: fileOut.Name(), Size: 0}, fileOut.Close()
}

func rmFile(path string) error {
	err := os.Remove(path)
	check(err)
	return err
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