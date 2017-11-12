package mediafile

import (
	"fmt"
	"videolib/logger"
)

var l = logger.SetLogger("mediaFile")

// TODO: path is valid

func CreateMedia(fullPath string) (*Video, *Subtitle, bool) {
	info := newBasicInfo(fullPath)
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
