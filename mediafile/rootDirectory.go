package mediafile

import (
	"videolib/fsutils"
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

	if fsutils.IsFile(path) {
		dirPath = fsutils.GetDirPath(path)
	} else {
		dirPath = fsutils.GetParentDirPath(path)
	}

	if dir, ok := r.DirMap[dirPath]; ok {
		return dir
	}

	parentDir := r.findDirByPathWithDependencies(dirPath)

	subDir := createEmtpyDir(dirPath)
	parentDir.Directories = append(parentDir.Directories, subDir)
	r.DirMap[dirPath] = subDir
	return subDir
}

func (r *RootDirectory) GetKeyPathsFromDirMap() []string {
	keysLength := len(r.DirMap)
	keys := make([]string, 0, keysLength)

	for k := range r.DirMap {
		keys = append(keys, k)
	}
	return keys
}

func (r *RootDirectory) GetDirectoryByKey(key string) (*Directory, bool) {
	dir, ok := r.DirMap[key]
	return dir, ok
}

func (r *RootDirectory) IsEmpty() bool {
	// firs is root
	return len(r.DirMap) <= 1
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
