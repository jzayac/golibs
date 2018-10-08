package dirscanner

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"videolib/logger"
)

func init() {
	logger.SetApp("dirScanner")
}

var l = logger.SetLogger("dirScanner")

// type FilterFileByTypeInterface interface {
// 	FilterFileByType(fileName string) bool
// }

func getFileType(name string) (string, bool) {
	index := strings.LastIndex(name, ".")
	if index == -1 {
		return "", false
	}

	return name[index+1:], true
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// type fileInformation struct {
// 	Name  string `json:"name"`
// 	IsDir string `json:"isDir"`
// }

func ListDirecotory(dir string) []os.FileInfo {
	cleanPath := path.Clean(dir)

	files, err := ioutil.ReadDir(cleanPath)

	if err != nil {
		return nil
	}

	return files
}

func FilterFilesInDirectoryByType(directory string, allowedTypes []string) []string {

	files := make([]string, 0, 100)

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		// fmt.Println(movie)
		fileType, haveFileType := getFileType(path)
		if !haveFileType {
			return nil
		}

		// if !fileTypes.FilterFileByType(fileType) {
		if !stringInSlice(fileType, allowedTypes) {
			return nil
		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		l.Error("something went wrong: "+err.Error(), "FilterFilesInDirectoryByType")
		// panic(err)
	}

	return files
}
