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

func sortFiles(sortMethod string, files []File) []File {
	var sortedFiles []File

	switch sortMethod {
	case "size asc":

	case "size desc":

	case "name A-Z":

	case "name Z-A":

	case "newest":

	case "oldest":

	case "modification (newest first)":

	case "modification (oldest first)":

	default:
		fmt.Println("something went wrong... sorting alphabetically by default")

	}

	return sortedFiles
}

func main() {
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

	//TODO: ask for sort method
	fmt.Println("How would you like to sort yoru files?")
	fmt.Println("1. Size (ascending)")
	fmt.Println("2. Size (descending")
	fmt.Println("3. Name (A-Z)")
	fmt.Println("4. Name (Z-A)")
	fmt.Println("5. Date Created (newest)")
	fmt.Println("6. Date Created (oldest)")
	fmt.Println("7. Date Modified (most recent)")
	fmt.Println("8. Date Modified (least recent)")

	var sortMethod string
	for {
		var shouldBreak = false
		//done: get input, prompt again if invalid
		//TODO: make this so you dont have to click enter
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
		}

		switch char {
		case '1':
			fmt.Println("1 pressed")
			sortMethod = "size asc"
			shouldBreak = true
			break
		case '2':
			fmt.Println("2 pressed")
			sortMethod = "size desc"
			shouldBreak = true
			break
		case '3':
			fmt.Println("3 pressed")
			sortMethod = "name A-Z"
			shouldBreak = true
			break
		case '4':
			fmt.Println("4 pressed")
			sortMethod = "name Z-A"
			shouldBreak = true
			break
		case '5':
			fmt.Println("5 pressed")
			sortMethod = "newest"
			shouldBreak = true
			break
		case '6':
			fmt.Println("6 pressed")
			sortMethod = "oldest"
			shouldBreak = true
			break
		case '7':
			fmt.Println("7 pressed")
			sortMethod = "modification (newest first)"
			shouldBreak = true
			break
		case '8':
			fmt.Println("8 pressed")
			sortMethod = "modification (oldest first)"
			shouldBreak = true
			break
		default:
			fmt.Println("Invalid input, please try again")
		}

		if shouldBreak {
			break
		}
	}

	//TODO: get array of files
	files := getFiles(rootDir, exts)

	//TODO: sort array
	files = sortFiles(sortMethod, files)

	//TODO: copy files to directory

}
