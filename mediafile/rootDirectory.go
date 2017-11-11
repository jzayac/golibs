package mediafile

import (
	"fmt"
	"videolib/dirutils"
)

// cannot have files in root dir.. only directories
type RootDirectory struct {
	directories []*Directory
	DirMap      map[string]*Directory
}

func (r *RootDirectory) findDirByPathWithDependencies(path string) *Directory {
	if r.directories == nil {
		l.Warn("Directories is empty", "findDirByPathWithDependencies")
		return nil
	}
	dirPath := ""

	fmt.Println(path)
	if dirutils.IsFile(path) {
		dirPath = dirutils.GetDirPath(path)
	} else {
		dirPath = dirutils.GetParentDirPath(path)
	}

	if dir, ok := r.DirMap[dirPath]; ok {
		// dir.Directories = append(dir.Directories, dir)
		return dir
	}

	parentDir := r.findDirByPathWithDependencies(dirPath)

	subDir := createEmtpyDir(dirPath)
	parentDir.Directories = append(parentDir.Directories, subDir)
	r.DirMap[dirPath] = subDir
	return parentDir
}

func NewRootDirectory() *RootDirectory {
	rootDir := createEmtpyDir("/")
	directories := make([]*Directory, 0, 20)
	directories = append(directories, rootDir)
	rootMap := make(map[string]*Directory)
	rootMap["/"] = rootDir
	return &RootDirectory{
		directories: directories,
		DirMap:      rootMap,
	}
}

func createEmtpyDir(path string) *Directory {
	videos := make([]*Video, 0, 20)
	subtitles := make([]*Subtitle, 0, 20)
	directories := make([]*Directory, 0, 20)
	return &Directory{
		Path:        path,
		Videos:      videos,
		Subtitles:   subtitles,
		Directories: directories,
	}
}
