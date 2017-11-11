package mediafile

// import (
// 	"testing"
// )

// // is this unit test?
// func TestCreateMediaTree(t *testing.T) {
// 	sliceOfPaths := []string{
// 		"/armagedon.mp4",
// 		"/movies/movie.mp4",
// 		"/movies/movie.srt",
// 		"/movies/armagedon.mp4",
// 		"/some/long/dir/deep.mp4",
// 	}

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
// }
