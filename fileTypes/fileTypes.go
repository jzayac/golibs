package fileTypes

var (
	SupportedVideoTypes = []string{
		"webm", "mkv", "flv", "ogv",
		"ogg", "avi", "mov", "qt", "wmv", "amv", "mp4", "m4p", "m4v",
		"mpg", "mp2", "mpeg", "mpe", "mpv", "mpg", "mpeg", "m2v",
	}
	webTypes               = []string{"mp4"}
	SupportedSubtitleTypes = []string{
		"sub", "srt", "stl", "ts", "ttml", "vtt",
	}
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IsVideoFileType(fileType string) bool {
	return stringInSlice(fileType, SupportedVideoTypes)
}
