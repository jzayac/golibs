package dirutils

import (
	"os"
	"strings"
)

func IsFile(path string) bool {
	sp := string(os.PathSeparator)
	if string(path[len(path)-1]) != sp {
		return true
	}
	return false
}

// directory string end with path separator
func GetDirPath(path string) string {
	sp := string(os.PathSeparator)
	if len(path) < 2 {
		// todo: should i log error if is empty string
		return sp
	}

	if IsFile(path) == false {
		return path
	}

	index := strings.LastIndex(path, sp)
	dirPath := path[:index+1]
	return dirPath
}

func GetDirName(path string) string {
	sp := string(os.PathSeparator)
	dirPath := GetDirPath(path)
	dirPath = dirPath[:len(dirPath)-1]

	if len(dirPath) < 2 {
		return sp
	}

	index := strings.LastIndex(dirPath, sp)
	dirName := dirPath[index+1:]

	return dirName
}

// directory string end with path separator
func GetParentDirPath(path string) string {
	sp := string(os.PathSeparator)
	if strings.Count(path, sp) < 2 {
		return sp
	}

	path = GetDirPath(path)

	// remove last path separator
	path = path[:len(path)-1]

	parrentDir := ""
	index := strings.LastIndex(path, sp)

	parrentDir = path[:index+1]

	return parrentDir
}

func GetParentDirName(path string) string {
	parrentPathDir := GetParentDirPath(path)
	return GetDirName(parrentPathDir)
}
