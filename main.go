package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const tmpDir = "/Users/taylorhoward"

type File struct {
	Info fs.FileInfo
	Path string
}

//TODO: make this retunr an array of the file type
//TODO: make this take in an array of strings for the extensions
func getFiles(path string, extensions []string) []File {
	var returnFiles []File

	files, _ := ioutil.ReadDir(path)
	for _, file := range files {

		fileExt := filepath.Ext(path + "/" + file.Name())

		fileObject := File{
			Info: file,
			Path: path + "/" + file.Name(),
		}
		for _, extension := range extensions {
			if fileExt == extension {
				//TODO: change this to add to slice, add from slice to folder in a different func
				//TODO: add file to folder
				//data, err := ioutil.ReadFile(path + "/" + file.Name())
				//if err != nil {
				//	fmt.Println(err)
				//	break
				//}
				//
				//err = ioutil.WriteFile(folderPath + "/" + file.Name(), data, 0777)
				//
				//if err != nil {
				//	fmt.Println(err)
				//	break
				//}
				//break

				returnFiles = append(returnFiles, fileObject)
				break
			}
		}
		//if fileExt == ".jpg" {
		//	fmt.Printf("Name: %v, size: %v kb, Mode: %v, isDir: %v\n",
		//		file.Name(),
		//		file.Size()/1000,
		//		file.Mode(),
		//		file.IsDir())
		//}

		if file.IsDir() {
			getFiles(path+"/"+file.Name(), extensions)
		}
	}

	return returnFiles
}

func sortFiles(sortMethod string, files map[fs.FileInfo]string) []string {
	var sortedFiles []string

	return sortedFiles
}

func main() {
	// var extensions [] string
	// var rootPath string
	// var newDirPath string
	reader := bufio.NewReader(os.Stdin)

	//Done: Ask for file type (ext)
	fmt.Println("Enter the file extensions to sort: <ext> <ext> <ext> etc.")
	args, _ := reader.ReadString('\n')
	args = strings.Replace(args, "\n", "", -1)

	fmt.Println(args)
	//done: parse input into slice
	exts := strings.Split(args, " ")

	_ = exts
	//TODO: if invalid, prompt again

	//done: Ask for root
	fmt.Println("Enter the root directory to sort from (starting with /)")
	fmt.Println("recommended starting point is '/Users/<youruser>'")
	var rootDir string
	for {
		rootDir, _ = reader.ReadString('\n')
		rootDir = strings.Replace(rootDir, "\n", "", -1)
		//if rootdir is directory, break
		//Done: if invalid, prompt again
		if file, err := os.Stat(rootDir); err == nil && file.IsDir() {
			break
		} else {
			fmt.Println(rootDir + " is not a valid directory, please enter a new root directory")
		}
	}

	//TODO: Ask for new folder location
	fmt.Println("enter the directory for the sorted folder")
	var folderDir string
	for {
		folderDir, _ = reader.ReadString('\n')
		folderDir = strings.Replace(folderDir, "\n", "", -1)

		if file, err := os.Stat(folderDir); err == nil && file.IsDir() {
			//TODO: create folder
			currentTime := time.Now()
			fmt.Println(currentTime.Format("01-02-2006 15:04:05"))
			fmt.Println(time.Now())
			if err := os.Mkdir(folderDir+"/FileSorter-"+args+"-"+currentTime.Format("01-01-2006 15:04:05"), 0777); err == nil {
				folderDir = folderDir + "/FileSorter-" + args + "-" + currentTime.Format("01-01-2006 15:04:05")
				break
			} else {
				//fmt.Println(err)
				fmt.Print("There was an error creating the folder: ")
				fmt.Println(err)
				fmt.Println("Please try again")
			}

		} else {
			fmt.Println(folderDir + " is not a valid directory, please enter a new location for the folder")
		}
	}

	//TODO: get array of files
	files := getFiles(rootDir, exts)
	_ = files

	//TODO: sort array

	//TODO: copy files to directory
}
