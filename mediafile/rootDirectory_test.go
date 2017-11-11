package mediafile

import (
	"fmt"
	"testing"
)

func TestFindDirByPathWithDependeciesDirectoryAddress(t *testing.T) {
	rootDirMock := &Directory{Path: "root"}
	subDirMock := &Directory{Path: "sub"}
	testDirMock := &Directory{Path: "test"}

	hashMapMock := make(map[string]*Directory)
	hashMapMock["/"] = rootDirMock
	hashMapMock["/sub/"] = subDirMock
	hashMapMock["/sub/test/"] = testDirMock
	rootDir := &RootDirectory{DirMap: hashMapMock}
	rootDir.directories = []*Directory{}

	cases := []struct {
		rootDir *RootDirectory
		path    string
		expect  *Directory
	}{
		{
			path:   "/sub/asdasd",
			expect: subDirMock,
		},
		{
			path:   "/sub/test/",
			expect: subDirMock,
		},
		{
			path:   "/sub/test/test.avi",
			expect: testDirMock,
		},
	}

	for _, c := range cases {
		foundDir := rootDir.findDirByPathWithDependencies(c.path)
		foundDirAddr := fmt.Sprintf("%p", foundDir)
		expectedDirAddr := fmt.Sprintf("%p", c.expect)

		if expectedDirAddr != foundDirAddr {
			t.Fatalf("EXPECT objects to be compare %s == %s", expectedDirAddr, foundDirAddr)
		}
	}
}

func TestReturnNilIfRootIsNotInitialized(t *testing.T) {
	// rootDir := NewRootDirectory()
	rootDir := &RootDirectory{}

	// test := &RootDirectory{}
	test := rootDir.findDirByPathWithDependencies("/")
	if test != nil {
		t.Fatalf("EXPECT to return nil")
	}
}

func TestFindDirByPathWithDependeciesRecusrsive(t *testing.T) {
	sliceOfPaths := []string{
		"/movies/armagedon/armagedon.mp4",
		"/movies/something/armagedon.mp4",
		"/movies/armagedon/armagedon.mp4",
		"/movies/movie.mp4",
		"/movies/movie.srt",
		"/movies/armagedon.mp4",
		"/some/long/dir/deep.mp4",
	}

	// 6 directories + root dir
	expectedLen := 7

	expect := []string{
		"/movies/armagedon/",
		"/movies/something/",
		"/movies/",
		"/some/long/dir/",
	}

	rootDir := NewRootDirectory()

	for _, filePath := range sliceOfPaths {
		rootDir.findDirByPathWithDependencies(filePath)
	}

	for _, path := range expect {
		_, ok := rootDir.DirMap[path]
		if !ok {
			t.Fatalf("Expect DirMap to have key %s", path)
		}
		// if
	}
	// t.Logf("%+v\n", rootDir.DirMap)

	if expectedLen != len(rootDir.DirMap) {
		t.Fatalf("Expect DirMap to have size %d, but have %d", expectedLen, len(rootDir.DirMap))
	}

	// 	dirTree := CreateMediaTree(sliceOfPaths)

	// 	if len(dirTree.Videos) != 1 {
	// 		t.Fatalf("length of Videos should be equal to 1")
	// 	}

	// 	if dirTree.Videos[0].Name != "armagedon" {
	// 		t.Fatalf("expected name for file in root directory is armagedon but get: %s", dirTree.Videos[0].Name)
	// 	}

	// 	if len(dirTree.Directories) != 2 {
	// 		t.Fatalf("length of Directories should be equal to 2")
	// 	}

	// 	if len(dirTree.Directories) != 2 {
	// 		t.Fatalf("length of Directories should be equal to 2")
	// 	}

	// 	if len(dirTree.Directories[0].Videos) != 2 {
	// 		t.Fatalf("length of Videos in subdirecory should be equal to 2")
	// 	}

	// 	if len(dirTree.Directories[0].Subtitles) != 1 {
	// 		t.Fatalf("length of Subtitles in subdirecory should be equal to 1")
	// 	}

	// 	if dirTree.Directories[0].Videos[0].Name != "movie" {
	// 		t.Fatalf("expected name for file in subdirectory is movie but get: %s", dirTree.Directories[0].Videos[0])
	// 	}

	// 	movieDir := dirTree.GetDirectoryByPath("/movies/")

	// 	if movieDir == nil {
	// 		t.Fatalf("directory /movies/ does not exist")
	// 	}

	// 	if len(movieDir.Videos) != 1 {
	// 		t.Fatalf("expect videos in  /movies/ dir is 1 but get %d", len(movieDir.Videos))
	// 	}

	// 	if dirTree.Directories[0].GetDirecoryName() != "movies" {
	// 		t.Fatalf("expected name for file in subdirectory is movies but get: %s", dirTree.Directories[0].GetDirecoryName())
	// 	}

	// 	if len(dirTree.Directories[1].Videos) != 0 {
	// 		t.Fatalf("length of Videos should be equal to 0")
	// 	}

	// 	if dirTree.Directories[1].Directories[0].Directories[0].Videos[0].Name != "deep" {
	// 		t.Fatalf("expected name for file in deep directory is deep but get: %s", dirTree.Directories[1].Directories[0].Directories[0].Videos[0].Name)
	// 	}

	// 	if dirTree.Directories[1].Directories[0].Directories[0].Path != "/some/long/dir" {
	// 		t.Fatalf("expected name for file in deep directory is deep but get: %s", dirTree.Directories[1].Directories[0].Directories[0].Path)
	// 	}

}
