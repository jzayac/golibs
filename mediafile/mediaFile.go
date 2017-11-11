package mediafile

import (
	"fmt"
	// "os"
	"videolib/logger"
)

var l = logger.SetLogger("mediaFile")

// fmt.Printf("%+v", l)

// TODO: path is valid

func CreateMedia(fullPath string) (*Video, *Subtitle, bool) {
	info := createBasicInfo(fullPath)
	if info == nil {
		return nil, nil, false
	}

	sub := NewSubtitle(*info)
	vid := NewVideo(*info)

	if sub == nil && vid == nil {
		return nil, nil, false
	}

	return vid, sub, true
}

func CreateSliceOfMediaFiles(sliceOfPaths []string) ([]*Video, []*Subtitle) {
	videos := make([]*Video, 0, 80)
	subtitles := make([]*Subtitle, 0, 80)

	for _, filePath := range sliceOfPaths {
		vid, sub, _ := CreateMedia(filePath)

		if vid != nil {
			videos = append(videos, vid)
		}

		if sub != nil {
			subtitles = append(subtitles, sub)
		}

	}
	return videos, subtitles
}

// func CreateMediaTree(sliceOfPaths []string) *Directory {
// 	sp := string(os.PathSeparator)
// 	mapTree := make(map[string]*Directory)
// 	mapTree[sp] = &Directory{}
// 	fmt.Printf("%+v", mapTree)

// 	for _, filePath := range sliceOfPaths {
// 		vid, sub, _ := CreateMedia(filePath)

// 		if vid != nil {
// 			dir := findParrendDir(mapTree, vid.Path)
// 			dir.Videos = append(dir.Videos, vid)
// 		}

// 		if sub != nil {
// 			dir := findParrendDir(mapTree, sub.Path)
// 			dir.Subtitles = append(dir.Subtitles, sub)
// 		}

// 	}

// 	fmt.Printf("%+v\n", mapTree)
// 	return mapTree[sp]
// }

// func GetMapedDirectories(sliceOfPaths []string) *RootDirectory {
func ParseListForMovies(sliceOfPaths []string) *RootDirectory {
	root := NewRootDirectory()

	for _, filePath := range sliceOfPaths {
		vid, sub, _ := CreateMedia(filePath)
		if vid != nil {
			dir := root.findDirByPathWithDependencies(vid.Path)
			dir.Videos = append(dir.Videos, vid)
		}

		if sub != nil {
			dir := root.findDirByPathWithDependencies(sub.Path)
			dir.Subtitles = append(dir.Subtitles, sub)
		}

	}

	l.Debug(fmt.Sprintf("%+v", root), "ParseListForMovies")
	return root
}

// func findParrendDir(mapTree map[string]*Directory, path string) *Directory {
// 	parent := dirutils.GetParentDirPath(path)
// 	// fmt.Println("findParentDir")
// 	// fmt.Println(path)
// 	// fmt.Println(parent)
// 	if dir, ok := mapTree[parent]; ok {
// 		return dir
// 	}
// 	parentDir := findParrendDir(mapTree, parent)

// 	videos := make([]*Video, 0, 20)
// 	subtitles := make([]*Subtitle, 0, 20)
// 	directories := make([]*Directory, 0, 20)
// 	subDir := &Directory{
// 		Path:        parent,
// 		Videos:      videos,
// 		Subtitles:   subtitles,
// 		Directories: directories,
// 	}
// 	parentDir.Directories = append(parentDir.Directories, subDir)
// 	mapTree[parent] = subDir
// 	return subDir
// }
