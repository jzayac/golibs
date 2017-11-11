package mediafile

import (
	"fmt"

	"videolib/dirutils"
)

type Directory struct {
	Path        string
	Directories []*Directory
	Subtitles   []*Subtitle
	Videos      []*Video
}

func (d Directory) GetDirectoryByPath(path string) *Directory {
	fmt.Println(path)

	return nil
}

func (d Directory) GetDirecoryName() string {
	fmt.Println(d.Path)
	fmt.Println(dirutils.GetParentDirName(d.Path))
	return dirutils.GetParentDirName(d.Path)
}

func (d Directory) HaveSubDirectory() bool {
	return len(d.Directories) > 0
}

func (d *Directory) SubDirHaveOnlySubtitles() bool {
	if !d.HaveSubDirectory() {
		return false
	}

	foundSubtitles := false
	for _, dir := range d.Directories {
		if len(dir.Videos) > 0 {
			return false
		}

		if len(dir.Subtitles) > 0 {
			foundSubtitles = true
		}
	}

	return foundSubtitles
}
