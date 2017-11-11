package mediafile

import (
	"testing"
)

func TestNewVideo(t *testing.T) {
	cases := []struct {
		BasicInfo
		Exist bool
	}{
		{
			BasicInfo: BasicInfo{FileType: "mp4"},
			Exist:     true,
		},
		{
			BasicInfo: BasicInfo{FileType: "non"},
			Exist:     false,
		},
	}

	for _, c := range cases {
		vid := NewVideo(c.BasicInfo)
		if (vid != nil) != c.Exist {
			t.Fatalf("Expected output to net exist")
		}
	}
}
