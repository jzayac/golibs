package mediafile

import (
	"os"
	"strings"
)

type Direcotry struct {
	Directories []*Direcotry
	Subtitles   []*Subtitle
	Videos      []*Video
}

func (d Direcotry) HaveSubDirectory() bool {
	return len(d.Directories) > 0
}

func (d *Direcotry) SubDirHaveOnlySubtitles() bool {
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

func getParentDirPath(path string) string {
	sp := string(os.PathSeparator)
	if strings.Count(path, sp) < 2 {
		return sp
	}
	parrentDir := ""
	index := strings.LastIndex(path, sp)

	parrentDir = path[:index]

	// index = strings.LastIndex(parrentDir, sp)
	// parrentDir = parrentDir[index:]

	return parrentDir
}
