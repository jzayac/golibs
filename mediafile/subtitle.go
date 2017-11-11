package mediafile

var supportedSubtitlesType = map[string]bool{
	"srt":  false,
	"stl":  false,
	"ts":   false,
	"ttml": false,
	"vtt":  true,
}

type Subtitle struct {
	BasicInfo
	Web bool
}

func NewSubtitle(info BasicInfo) *Subtitle {
	fileType := info.FileType
	web, subs := supportedSubtitlesType[fileType]

	if !subs {
		return nil
	}

	sub := &Subtitle{BasicInfo: info, Web: web}
	return sub
}
