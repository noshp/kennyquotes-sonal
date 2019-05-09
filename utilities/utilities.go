package utilities

import (
	"os"
	"path/filepath"
)

//Check error function that can be used in other packages
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

//FilePathWalkDir walks through a list of dirs, if its not a file walk through again returns a slice of files found
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
