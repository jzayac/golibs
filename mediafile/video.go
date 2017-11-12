package mediafile

var supportedVideoType = map[string]bool{
	"mp4": true,
	"avi": false,
	"wmv": true,
}

type Video struct {
	BasicInfo
	Web bool
}

// interface for videoPlayer
// need to refactor
func (v Video) ParsedName() string {
	// return v.GetName()
	return parseName(v.Name)
}

// interface for videoPlayer
// need to refactor
func (v Video) ParseYearFromName() string {
	return v.GetYear()
}

func (v Video) GetName() string {
	return v.Name
}

func (v Video) GetVideoDbOptions() map[string]string {
	year := v.GetYear()
	return map[string]string{
		"include_adult":        "false",
		"year":                 year,
		"orimary_release_year": year,
	}
}

func (v Video) ParseNameByDirName() (string, bool) {
	str, found := getNameFromFolderName(v.GetDirName())
	return str, found
}

func NewVideo(info BasicInfo) *Video {
	fileType := info.FileType
	web, val := supportedVideoType[fileType]

	if !val {
		return nil
	}

	video := &Video{BasicInfo: info, Web: web}
	return video
}
