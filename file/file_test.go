package file

import (
	"testing"
)

func TestIsDir(t *testing.T) {
	cases := []struct {
		file   *fileUtil
		expect bool
	}{
		{
			file: &fileUtil{
				path: "/test/movie1.mp4",
			},
			expect: false,
		},
		{
			file: &fileUtil{
				path: "/test/",
			},
			expect: true,
		},
		{
			file: &fileUtil{
				path: "/test/movie",
			},
			expect: false,
		},

		{
			file: &fileUtil{
				path: "",
			},
			expect: true,
		},
	}

	for _, c := range cases {
		dir := c.file.isDir()

		if dir != c.expect {
			t.Errorf("isdir invalid %s return %t", c.file.path, dir)
		}
	}

}

func TestGetFileNameFromPath(t *testing.T) {
	cases := []struct {
		file   *fileUtil
		expect string
	}{
		{
			file: &fileUtil{
				path: "/test/movie1.mp4",
			},
			expect: "movie1.mp4",
		},
		{
			file: &fileUtil{
				path: "/test/movie.avi",
			},
			expect: "movie.avi",
		},
		{
			file: &fileUtil{
				path: "/test/movie.ogg",
			},
			expect: "movie.ogg",
		},
		{
			file: &fileUtil{
				path: "/test/",
			},
			expect: "",
		},
		{
			file: &fileUtil{
				path: "/test/movie.rit",
			},
			expect: "movie.rit",
		},
		{
			file: &fileUtil{
				path: "/test/movie.avi",
			},
			expect: "movie.avi",
		},
	}

	for _, c := range cases {
		path := c.file.GetFileNameFromPath()

		if path != c.expect {
			t.Errorf("unexpected filen name %s but get %s", c.expect, path)
		}
	}

}

func TestGetFileType(t *testing.T) {
	cases := []struct {
		file   *fileUtil
		expect string
	}{
		{
			file: &fileUtil{
				path: "/test/movie1.mp4",
			},
			expect: "mp4",
		},
		{
			file: &fileUtil{
				path: "/test/movie.avi",
			},
			expect: "avi",
		},
		{
			file: &fileUtil{
				path: "/test/movie.ogg",
			},
			expect: "ogg",
		},
		{
			file: &fileUtil{
				path: "/test/",
			},
			expect: "",
		},
		{
			file: &fileUtil{
				path: "/test/movie.rit",
			},
			expect: "rit",
		},
		{
			file: &fileUtil{
				path: "/test/mo.vie.avi",
			},
			expect: "avi",
		},
		{
			file: &fileUtil{
				path: "/te.st/mo.vie.avi",
			},
			expect: "avi",
		},
	}

	for _, c := range cases {
		path := c.file.GetFileType()

		if path != c.expect {
			t.Errorf("unexpected filen type %s but get %s", c.expect, path)
		}
	}

}

func TestParseMovstringieName(t *testing.T) {
	cases := []struct {
		file        *fileUtil
		expectTitle string
		expectYear  int
	}{
		{
			file: &fileUtil{
				path: "/files/other/Titanic (1997)/Titanic.1997.720p.HDTV.x264-YIFY.mp4",
			},
			expectTitle: "Titanic",
			expectYear:  1997,
			// expect: "mp4",
		},
		{
			file: &fileUtil{
				path: "/files/movie/A.Nightmare.On.Elm.Street.3.Dream.Warriors.1987.1080p.BluRay.H264.AAC-RARBG/A.Nightmare.On.Elm.Street.3.Dream.Warriors.1987.1080p.BluRay.H264.AAC-RARBG.mp4",
			},

			expectTitle: "A Nightmare On Elm Street 3 Dream Warriors",
			expectYear:  1987,
		},
		{
			file: &fileUtil{
				path: "/files/movie/Demolition Man (1993)/Demolition.Man.1993.720p.BRrip.x264.YIFY.mp4",
			},
			expectTitle: "Demolition Man",
			expectYear:  1993,
		},
		{
			file: &fileUtil{
				path: "/files/movie/Dead.Snow.2.Red.vs.Dead.(Doed.Snoe.2).2014.1080p.BluRay.x264.anoXmous/Dead.Snow.2.Red.vs.Dead.(Doed.Snoe.2).2014.1080p.BluRay.x264.anoXmous_.mp4",
			},
			expectTitle: "Dead Snow 2 Red vs Dead",
			expectYear:  2014,
		},
		{
			file: &fileUtil{
				path: "/files/movie/Dech-života-(1987;-televizní-dabing).mp4",
			},
			expectTitle: "Dech-života-",
			expectYear:  1987,
		},
	}

	// t.Logf("%+v\n", info)
	for _, c := range cases {
		info, err := c.file.ParseName()

		// log.Printf("%+v\n", info)
		// t.Errorf("%+v\n", info)
		if err != nil {
			t.Errorf("Somethnig went wrong %s", err)
		}

		if info.Title != c.expectTitle {
			t.Errorf("wrong parsed file %s get %s", c.expectTitle, info.Title)
		}

		if info.Year != c.expectYear {
			t.Errorf("wrong parsed file %d get %d", c.expectYear, info.Year)
		}

	}
}

func TestParseTvshowName(t *testing.T) {
	cases := []struct {
		file          *fileUtil
		expectTitle   string
		expectSeason  int
		expectEpisode int
	}{
		{
			file: &fileUtil{
				path: "/files/tvshow/Akta X/Season 02/02x04 Nespaví.avi",
			},
			expectTitle:   "Akta X",
			expectSeason:  2,
			expectEpisode: 4,
		},
		{
			file: &fileUtil{
				path: "/files/tvshow/Frasier/Season 01/01x11.avi",
			},
			expectTitle:   "Frasier",
			expectSeason:  1,
			expectEpisode: 11,
		},
		{
			file: &fileUtil{
				path: "/files/tvshow/The IT Crowd/Season 1/01-Yesterday's Jam.avi",
			},
			expectTitle:   "The IT Crowd",
			expectSeason:  0,
			expectEpisode: 0,
		},
	}

	for _, c := range cases {
		info, err := c.file.ParseNameFromDir()

		if err != nil {
			t.Errorf("Somethnig went wrong %s", err)
		}

		if info.Title != c.expectTitle {
			t.Errorf("wrong title %s get %s", c.expectTitle, info.Title)
		}

		if info.Season != c.expectSeason {
			t.Errorf("wrong season %d get %d", c.expectSeason, info.Season)
		}
		if info.Episode != c.expectEpisode {
			t.Errorf("wrong episode %d get %d", c.expectEpisode, info.Episode)
		}
	}
}
