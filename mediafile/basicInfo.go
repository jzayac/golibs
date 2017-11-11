package mediafile

import (
	"os"
	"strings"

	"videolib/dirutils"
)

type BasicInfo struct {
	Name     string
	Path     string
	FileType string
}

func (v BasicInfo) GetDirName() string {
	return dirutils.GetDirName(v.Path)
}

func (v *BasicInfo) ParseParentDirName() (string, bool) {
	return getNameFromFolderName(v.GetDirName())
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
