package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const tmpDir = "/Users/taylorhoward"

//TODO: make this retunr an array of the file type
//TODO: make this take in an array of strings for the extensions
func getFiles(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {

		ext := filepath.Ext(path + "/" + file.Name())

		if ext == ".jpg" {
			fmt.Printf("Name: %v, size: %v kb, Mode: %v, isDir: %v\n",
				file.Name(),
				file.Size()/1000,
				file.Mode(),
				file.IsDir())
		}

		if file.IsDir() {
			getFiles(path + "/" + file.Name())
		}
	}
}

func main() {
	//TODO: Ask for file type (ext)
	//TODO: Ask for root
	//TODO: Ask for new folder location

	//TODO: mkdir

	//TODO: get array of files
	getFiles(tmpDir)

	//TODO: sort array

	//TODO: copy files to directory
}
