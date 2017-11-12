package mediafile

type Directory struct {
	Path        string
	Directories []*Directory
	Subtitles   []*Subtitle
	Videos      []*Video
}

// func (d Directory) GetDirectoryByPath(path string) *Directory {
// 	fmt.Println(path)
// 	return nil
// }

// func (d Directory) GetParentDirecoryName() string {
// 	return fsutils.GetParentDirName(d.Path)
// }

func (d Directory) VideoFileIsAlone() bool {
	onlyOneVideo := d.OnlyOneVideoInDirectory()
	haveSubdir := d.HaveSubDirectory()

	if !onlyOneVideo {
		return false
	}

	if !haveSubdir {
		return true
	}

	if d.SubDirHaveOnlySubtitles() {
		return true
	}

	return false
}

func (d Directory) OnlyOneVideoInDirectory() bool {
	return len(d.Videos) == 1
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
