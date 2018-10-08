package mediafile

import "testing"

func TestGetClearName(t *testing.T) {
	cases := []struct {
		name   string
		expect string
	}{
		{
			name:   "Wonder.Woman.2017.TC1080P.x264",
			expect: "wonder woman",
		},
		{
			name:   "gothika (2003)",
			expect: "gothika",
		},
		{
			name:   "The Avengers (2012) 1080p_6CH_BRrip_scOrp_sujaidr",
			expect: "the avengers",
		},
		{
			name:   "Hellboy The Golden Army (2008)",
			expect: "hellboy the golden army",
		},
		// {
		// 	name:   "star-trek-vii-generace-cz-iculdo",
		// 	expect: "star trek vii generace",
		// },
		{
			name:   "Frozen.2013.480p.BDRip.XviD.CZ",
			expect: "frozen",
		},
		{
			name:   "Fantomas-kontra-Scotland-Yard",
			expect: "fantomas kontra scotland yard",
		},
		{
			name:   "Vladce-temnot_[horor-1987]-CZ-dab",
			expect: "vladce temnot",
		},
		{
			name:   "Pirates.of.the.Caribbean.Dead.Men.Tell.No.Tales.2017.HDRip.XviD.mp4",
			expect: "pirates of the caribbean dead men tell no tales",
		},
		{
			name:   "Batman [2008] The Dark Knight hd",
			expect: "batman",
		},
		{
			name:   "Batman Dark Knight Rises 720p 2012",
			expect: "batman dark knight rises",
		},
		{
			name:   "[2008] Batman The Dark Knight",
			expect: "batman the dark knight",
		},
	}

	for _, c := range cases {
		clearName := parseName(c.name)
		if clearName != c.expect {
			t.Fatalf("Expected name to be %s but it was %s", c.expect, clearName)

		}
	}
}

func TestGetNameFromFolderName(t *testing.T) {
	cases := []struct {
		path   string
		expect string
	}{
		{
			path:   "Avatar (2009)",
			expect: "avatar",
		},
		{
			path:   "he Usual Suspects 1995",
			expect: "",
		},
		{
			path:   "Pulp Fiction (1994)",
			expect: "pulp fiction",
		},
		{
			path:   "Reservoir Dogs (1995)",
			expect: "reservoir dogs",
		},
		{
			path:   "Reservoir Dogs 1995",
			expect: "",
		},
	}

	for _, c := range cases {
		info, ok := getNameFromFolderName(c.path)
		if ok == (len(c.expect) == 0) {
			t.Fatalf("expect to return empty string but get %s", info)
		}
		if info != c.expect {
			t.Fatalf("expect to be value %s but get %s", c.expect, info)
		}
	}
}

func TestFindYearIndexFromString(t *testing.T) {
	cases := []struct {
		name   string
		expect int
	}{
		{
			name:   "Blade Runner 2049 2019",
			expect: -1,
		},
		{
			name:   "Blade Runner 2049 2017",
			expect: 18,
		},
		{
			name:   "Blade Runner (2049) 2017",
			expect: 20,
		},
		{
			name:   "Blade Runner 2017 2017",
			expect: 18,
		},
		{
			name:   "Blade Runner (2017) 2017",
			expect: 14,
		},
		{
			name:   "Pirates.of.the.Caribbean.Dead.Men.Tell.No.Tales.2017.HDRip.XviD.",
			expect: 48,
		},
		{
			name:   "Reservoir Dogs (1995)",
			expect: 16,
		},
	}

	for _, c := range cases {
		index := findYearIndexFromString(c.name, false)
		if index != c.expect {

			t.Fatalf("expect index to be value %d but get %d", c.expect, index)
		}
	}
}

func TestRegexYearString(t *testing.T) {
	str := regexYearString(2017)
	if str != `(20([0]\d{1}|1[0-7]))` {
		t.Fatalf("wrong regex %s", str)
	}
	str = regexYearString(2019)
	if str != `(20([0]\d{1}|1[0-9]))` {
		t.Fatalf("wrong regex %s", str)
	}
}
