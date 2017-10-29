package mediafile

import (
	"testing"
)

func TestCreateSubtitle(t *testing.T) {
	cases := []struct {
		BasicInfo
		Exist bool
	}{
		{
			BasicInfo: BasicInfo{FileType: "srt"},
			Exist:     true,
		},
		{
			BasicInfo: BasicInfo{FileType: "non"},
			Exist:     false,
		},
		{
			BasicInfo: BasicInfo{FileType: "ts"},
			Exist:     true,
		},
	}

	for _, c := range cases {
		vid := CreateSubtitle(c.BasicInfo)
		if (vid != nil) != c.Exist {
			t.Fatalf("Expected output to net exist")
		}
	}
}
