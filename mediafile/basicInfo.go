package mediafile

import (
	"fmt"
	"os"
	"strings"

	"videolib/fsutils"
)

type BasicInfo struct {
	Name     string
	Path     string
	FileType string
}

func (v BasicInfo) GetPath() string {
	return v.Path
}

func (v BasicInfo) GetFileType() string {
	return v.FileType
}

func (v BasicInfo) GetDirName() string {
	return fsutils.GetDirName(v.Path)
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

func newBasicInfo(fullPath string) *BasicInfo {
	sp := string(os.PathSeparator)
	fmt.Println("KOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO1")
	index := strings.LastIndex(fullPath, sp)
	name := ""

	if index != -1 {
		name = fullPath[index+1:]
	}

	fmt.Println("KOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO2")
	index = strings.LastIndex(name, ".")
	if index == -1 {
		return nil
	}

	fmt.Println("KOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO3")
	info := &BasicInfo{}
	info.Name = name[0:index]
	info.FileType = name[index+1:]
	info.Path = fullPath

	fmt.Println("KOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	fmt.Printf("%s\n", "CHUJUUUUUUUUUUUUUUUUU")
	fmt.Printf("%s\n", "picaaaaaaaaaaaaaaaaaaaaa")
	fmt.Printf("%+v\n", info)
	fmt.Printf("%+v\n", info)
	fmt.Printf("%+v\n", info)

	fmt.Println("KOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	return info
}
