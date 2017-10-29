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
func (v Video) ParseName() string {
	return v.GetName()
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

func CreateVideo(info BasicInfo) *Video {
	fileType := info.FileType
	web, val := supportedVideoType[fileType]

	if !val {
		return nil
	}

	video := &Video{BasicInfo: info, Web: web}
	return video
}
