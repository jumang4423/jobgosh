package utils

import (
	"fmt"
	"io/ioutil"
	types "jobgosh/src/typesAndVars"
	"os"
	"path/filepath"
)

func Dirsearch(dir string) []string {

	// get file path then return the array of file paths
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(string(types.ColorRed), "! error #03, log directory error", string(types.ColorReset))
		os.Exit(0)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, Dirsearch(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
