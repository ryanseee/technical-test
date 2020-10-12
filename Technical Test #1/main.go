package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := getCurrentDir()
	files := getAllFiles(dir)
	filesWithTODO := getFilesWithTODO(files)

	for _, file := range filesWithTODO {
		fmt.Println(file)
	}

	// var filesWithTODO []string

	// for _, file := range files {
	// 	data, err := ioutil.ReadFile(file)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fileItems := string(data)
	// 	line := 0
	// 	temp := strings.Split(fileItems, "\n")

	// 	for _, item := range temp {
	// 		if strings.Contains(item, "TODO") {
	// 			filesWithTODO = append(filesWithTODO, file)
	// 			break
	// 		}
	// 		line++
	// 	}

	// }
}

/** Get current directory of where main.go is residing **/
func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)
	return dir
}

/** Walks through all folders and finds all files only (without dir/sub-dir folder names) **/
func getAllFiles(dir string) []string {
	var files []string

	root := dir
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files

}

/** Get files with any line that contains TODO **/
func getFilesWithTODO(files []string) []string {
	var filesWithTODO []string

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		fileItems := string(data)
		line := 0
		temp := strings.Split(fileItems, "\n")

		for _, item := range temp {
			if strings.Contains(item, "TODO") {
				filesWithTODO = append(filesWithTODO, file)
				break
			}
			line++
		}

	}

	return filesWithTODO
}
