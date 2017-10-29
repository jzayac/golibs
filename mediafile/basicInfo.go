package mediafile

import (
	"os"
	// "regexp"
	"strings"
)

type BasicInfo struct {
	Name     string
	Path     string
	FileType string
}

func (v BasicInfo) GetParrentDirName() string {
	sp := string(os.PathSeparator)

	parrentPathDir := getParentDirPath(v.Path)

	if strings.Count(v.Path, sp) < 2 {
		return sp
	}
	index := strings.LastIndex(parrentPathDir, sp)
	parrentDirName := parrentPathDir[index+1:]

	return parrentDirName
}

func (v *BasicInfo) ParseParentDirName() (string, bool) {
	return getNameFromFolderName(v.GetParrentDirName())
}

func (v BasicInfo) GetYear() string {
	idx := findYearIndexFromString(v.Name, false)
	if idx == -1 {
		return ""
	}
	return v.Name[idx : idx+4]
}

func createBasicInfo(fullPath string) *BasicInfo {
	sp := string(os.PathSeparator)
	index := strings.LastIndex(fullPath, sp)
	name := ""

	if index != -1 {
		name = fullPath[index+1:]
	}

	index = strings.LastIndex(name, ".")
	if index == -1 {
		return nil
	}

	info := &BasicInfo{}
	info.Name = name[0:index]
	info.FileType = name[index+1:]
	info.Path = fullPath

	return info
}
