package mediafile

import (
	"reflect"
	"testing"
)

func TestCreateBasicInfo(t *testing.T) {
	cases := []struct {
		path string
		BasicInfo
	}{
		{
			path:      "/test/file.txt",
			BasicInfo: BasicInfo{Name: "file", FileType: "txt", Path: "/test/file.txt"},
		},
		{
			path:      "/test2/file2.none",
			BasicInfo: BasicInfo{Name: "file2", FileType: "none", Path: "/test2/file2.none"},
		},
	}

	for _, c := range cases {
		info := createBasicInfo(c.path)
		if !reflect.DeepEqual(*info, c.BasicInfo) {
			t.Fatalf("Expected to be equal %+v :: %+v", info, c.BasicInfo)
		}
	}
}

func TestGetYear(t *testing.T) {
	cases := []struct {
		BasicInfo
		expect string
	}{
		{
			BasicInfo: BasicInfo{Name: "Wonder.Woman.2017.TC1080P.x264"},
			expect:    "2017",
		},
		{
			BasicInfo: BasicInfo{Name: "gothika (2003)"},
			expect:    "2003",
		},
		{
			BasicInfo: BasicInfo{Name: "The Avengers (2012) 1080p_6CH_BRrip_scOrp_sujaidr"},
			expect:    "2012",
		},
		{
			BasicInfo: BasicInfo{Name: "Hellboy The Golden Army (2008)"},
			expect:    "2008",
		},
		{
			BasicInfo: BasicInfo{Name: "star-trek-vii-generace-cz-iculdo"},
			expect:    "",
		},
		{
			BasicInfo: BasicInfo{Name: "Frozen.2013.480p.BDRip.XviD.CZ"},
			expect:    "2013",
		},
		{
			BasicInfo: BasicInfo{Name: "Vladce-temnot_[horor-1987]-CZ-dab"},
			expect:    "1987",
		},
	}

	for _, c := range cases {
		year := c.GetYear()
		if year != c.expect {
			t.Fatalf("Expected name to be %s but it was %s", c.expect, year)

		}
	}
}

func TestGetDirName(t *testing.T) {
	cases := []struct {
		BasicInfo
		expect string
	}{
		{
			BasicInfo: BasicInfo{Path: "/movie/test1/something"},
			expect:    "test1",
		},
		{
			BasicInfo: BasicInfo{Path: "/"},
			expect:    "/",
		},
		{
			BasicInfo: BasicInfo{Path: ""},
			expect:    "/",
		},
	}

	for _, c := range cases {
		name := c.GetDirName()
		if name != c.expect {
			t.Fatalf("Expected name to be %s but it was %s", c.expect, name)

		}
	}
}
