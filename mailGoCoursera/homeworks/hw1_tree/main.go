package main

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
)

func dirWithoutFiles(files []fs.FileInfo) []fs.FileInfo {
	dwf := make([]fs.FileInfo, 0)
	for _, file := range files {
		if file.IsDir() {
			dwf = append(dwf, file)
		}
	}
	return dwf
}

func dirSubTree(out io.Writer, path string, printFiles bool, levelPrefix string) (err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	if !printFiles {
		files = dirWithoutFiles(files)
	}
	prefix := "├───"
	for idx, file := range files {
		if idx == len(files)-1 {
			prefix = "└───"
		}
		var filesize string
		if printFiles && !file.IsDir() {
			filesize = " (empty)"
			if file.Size() != 0 {
				filesize = " (" + strconv.Itoa(int(file.Size())) + "b)"
			}
		}
		fmt.Fprint(out, levelPrefix, prefix, file.Name(), filesize, "\n")
		if file.IsDir() {
			var subPath, subLevelPrefix string
			if path == "." {
				subPath = file.Name()
			} else {
				subPath = path + string(os.PathSeparator) + file.Name()
			}
			if idx != len(files)-1 {
				subLevelPrefix += "│"
			}
			dirSubTree(out, subPath, printFiles, levelPrefix+subLevelPrefix+"\t")
		}
	}
	return
}

func dirTree(out io.Writer, path string, printFiles bool) (err error) {
	return dirSubTree(out, path, printFiles, "")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-f] path\n", os.Args[0])
		os.Exit(1)
	}
	path := os.Args[1]
	printFiles := len(os.Args) > 2 && os.Args[2] == "-f"
	err := dirTree(os.Stdout, path, printFiles)
	if err != nil {
		panic(err)
	}
}
