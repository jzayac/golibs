package file

import "path/filepath"
import PTN "github.com/middelink/go-parse-torrent-name"
import "errors"
import "strings"

type fileUtil struct {
	path string
}

func (f fileUtil) isDir() bool {
	if len(f.path) == 0 {
		// ??
		return true
	}

	if f.path[len(f.path)-1] == filepath.Separator {
		return true
	}
	return false
}

func (f fileUtil) GetFileNameFromPath() string {
	if f.isDir() {
		return ""
	}
	file := filepath.Base(f.path)
	return file
}

func (f fileUtil) GetFileType() string {
	if f.isDir() {
		return ""
	}
	ext := filepath.Ext(f.path)
	if len(ext) <= 1 {
		return ""
	}
	return ext[1:]
}

func (f fileUtil) GetFullPath() string {
	return f.path
}

func (f fileUtil) ParseName() (*PTN.TorrentInfo, error) {
	name := f.GetFileNameFromPath()
	if name == "" {
		return nil, errors.New("NOT ABLE TO PARSE")
	}
	info, err := PTN.Parse(name)
	return info, err
}

// https://support.plex.tv/articles/naming-and-organizing-your-tv-show-files/
// file itsefs must contain season num and episode num
// name of tvshow must be parsed from dir
// for example /tvshow/Grey's Anatomy/Season 02/s02e01.avi
func (f fileUtil) ParseNameFromDir() (*PTN.TorrentInfo, error) {
	name := f.GetFileNameFromPath()
	info, err := PTN.Parse(name)

	dir := filepath.Dir(f.path)
	dirName := filepath.Base(dir)

	idx := strings.Index(strings.ToLower(dir), "season")
	if idx == -1 {
		info.Title = dirName
		return info, nil
	}

	dir = filepath.Dir(dir)
	dirName = filepath.Base(dir)
	info.Title = dirName

	return info, err
}

func NewFile(path string) *fileUtil {
	return &fileUtil{
		// path: filepath.Clean(path),
		path: path,
	}
}
