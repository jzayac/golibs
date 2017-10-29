package mediafile

import (
	"os"
)

// TODO: CREATE LOG LIBRARY

func CreateMedia(fullPath string) (*Video, *Subtitle, bool) {
	info := createBasicInfo(fullPath)
	if info == nil {
		return nil, nil, false
	}

	sub := CreateSubtitle(*info)
	vid := CreateVideo(*info)

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

func CreateMediaTree(sliceOfPaths []string) *Direcotry {
	sp := string(os.PathSeparator)
	mapTree := make(map[string]*Direcotry)
	mapTree[sp] = &Direcotry{}

	for _, filePath := range sliceOfPaths {
		vid, sub, _ := CreateMedia(filePath)

		if vid != nil {
			dir := findParrendDir(mapTree, vid.Path)
			dir.Videos = append(dir.Videos, vid)
		}

		if sub != nil {
			dir := findParrendDir(mapTree, sub.Path)
			dir.Subtitles = append(dir.Subtitles, sub)
		}

	}

	return mapTree[sp]
}

func findParrendDir(mapTree map[string]*Direcotry, path string) *Direcotry {
	parent := getParentDirPath(path)
	if dir, ok := mapTree[parent]; ok {
		return dir
	}
	parentDir := findParrendDir(mapTree, parent)

	videos := make([]*Video, 0, 20)
	subtitles := make([]*Subtitle, 0, 20)
	directories := make([]*Direcotry, 0, 20)
	subDir := &Direcotry{
		Videos:      videos,
		Subtitles:   subtitles,
		Directories: directories,
	}
	parentDir.Directories = append(parentDir.Directories, subDir)
	mapTree[parent] = subDir
	return subDir
}
