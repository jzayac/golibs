package fsutils

import (
	"testing"
)

func TestIsFile(t *testing.T) {
	cases := []struct {
		path   string
		expect bool
	}{
		{
			path:   "/movies/test/asdasd.123",
			expect: true,
		},
		{
			path:   "/movies/test/sub",
			expect: true,
		},
		{
			path:   "/movies/test/",
			expect: false,
		},
		{
			path:   "/",
			expect: false,
		},
	}

	for _, c := range cases {
		exc := IsFile(c.path)
		if exc != c.expect {
			t.Fatalf("Expect to not be %t with file %s", exc, c.path)
		}
	}
}

func TestGetDirName(t *testing.T) {
	cases := []struct {
		path   string
		expect string
	}{
		{
			path:   "/movies/test.rit",
			expect: "movies",
		},
		{
			path:   "/movies/test/sub/",
			expect: "sub",
		},
		{
			path:   "/movies/test/sub/more",
			expect: "sub",
		},
		{
			path:   "/",
			expect: "/",
		},
	}

	for _, c := range cases {
		name := GetDirName(c.path)
		if name != c.expect {
			t.Fatalf("Expected name is %s but get %s", c.expect, name)
		}
	}
}

func TestGetParentDirName(t *testing.T) {
	cases := []struct {
		path   string
		expect string
	}{
		{
			path:   "/movies/test.rit",
			expect: "/",
		},
		{
			path:   "/movies/test/sub/",
			expect: "test",
		},
		{
			path:   "/movies/test/sub/more",
			expect: "test",
		},
	}

	for _, c := range cases {
		name := GetParentDirName(c.path)
		if name != c.expect {
			t.Fatalf("Expected name is %s but get %s", c.expect, name)
		}
	}
}

func TestGetParentDirPath(t *testing.T) {
	cases := []struct {
		path   string
		expect string
	}{
		{
			path:   "/movies/test/asdasd.123",
			expect: "/movies/",
		},
		{
			path:   "/movies/test/sub",
			expect: "/movies/",
		},
		{
			path:   "/series/test/",
			expect: "/series/",
		},
		{
			path:   "/series/test",
			expect: "/",
		},
		{
			path:   "/movies/",
			expect: "/",
		},
		{
			path:   "/",
			expect: "/",
		},
		{
			path:   "/movies/test/sub/",
			expect: "/movies/test/",
		},
	}

	for _, c := range cases {
		name := GetParentDirPath(c.path)
		if name != c.expect {
			t.Fatalf("Expected name is %s but get %s", c.expect, name)
		}
	}
}

func TestGetDirPath(t *testing.T) {
	cases := []struct {
		path   string
		expect string
	}{
		{
			path:   "/movies/test/asdasd.123",
			expect: "/movies/test/",
		},
		{
			path:   "/movies/test/sub",
			expect: "/movies/test/",
		},
		{
			path:   "/movies/test/",
			expect: "/movies/test/",
		},
		{
			path:   "/",
			expect: "/",
		},
	}

	for _, c := range cases {
		exc := GetDirPath(c.path)
		if exc != c.expect {
			t.Fatalf("Expect path to be %s but get %s", exc, c.expect)
		}
	}
}
