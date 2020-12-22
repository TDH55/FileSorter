package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

const tmpDir = "/Users/taylorhoward"

//TODO: make this retunr an array of the file type
//TODO: make this take in an array of strings for the extensions
func getFiles(path string) []fs.FileInfo {
	var returnFiles []fs.FileInfo

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

	return returnFiles
}

func main() {
	// var extensions [] string
	// var rootPath string
	// var newDirPath string

	//TODO: Ask for file type (ext)
	fmt.Println("Enter the file extensions to sort: <ext> <ext> <ext> etc.")
	var arguments string
	fmt.Scanln(&arguments)
	//TODO: parse input into slice

	//TODO: if invalid, prompt again

	//TODO: Ask for root
	fmt.Println("Enter the root directory to sort from (starting with /)")
	fmt.Println("reccommended starting point is '/Users/<youruser>'")
	//TODO: if invalid, prompt again

	//TODO: Ask for new folder location
	fmt.Println("enter the directory for the sorted foler")
	//TODO: mkdir, if fails, prompt again

	//TODO: get array of files
	getFiles(tmpDir)

	//TODO: sort array

	//TODO: copy files to directory
}
