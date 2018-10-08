package mediafile

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// func parseNameOld(name string) string {
// 	// ak bude rok jebnit to vsetko do toho roku ak nie obmedzit to na 4 slova
// 	re := regexp.MustCompile(`(<[\w|\d|\s]*>)|(\[[\w|\d|\s]*\])|(\d{4}.+)`)
// 	result := re.ReplaceAllLiteralString(name, "")
// 	fmt.Println(result)
// 	result = strings.Replace(result, ".", " ", -1)
// 	result = strings.Replace(result, "-", " ", -1)
// 	re = regexp.MustCompile(`^([a-zA-Z]+\s*){1,4}`)
// 	// re = regexp.MustCompile(`^([a-zA-Z]+\s*){1,}`)
// 	result = re.FindString(result)
// 	result = strings.ToLower(result)
// 	return strings.Trim(result, " ")
// }

// TODO: also number in movie
// http://kodi.wiki/view/Naming_video_files/Movies
// https://support.plex.tv/hc/en-us/articles/200288586-Installation
func parseName(name string) string {
	// yearFound := false

	// fmt.Println("ORIGINAL NAME " + name)
	name = strings.ToLower(name)
	if idx := findYearIndexFromString(name, false); idx != -1 {
		// yearFound = true
		if idx > 4 {
			name = name[:idx-1]
		}
	}

	if idx := strings.Index(name, "bluray"); idx != -1 {
		name = name[:idx]
	}

	re := regexp.MustCompile(`(<.*>)|(\(.*\))|(\[.*\])`)
	result := re.ReplaceAllLiteralString(name, "")
	result = strings.Replace(result, ".", " ", -1)
	result = strings.Replace(result, "-", " ", -1)

	result = strings.Replace(result, "720p", " ", -1)
	result = strings.Replace(result, "1080p", " ", -1)

	// fmt.Println("REPLACE " + result)

	// regexStr := `^([a-zA-Z0-9][^720p|1080p]+\s*)`
	regexStr := `([a-z0-9]+\s*)`
	// regexStr := `^([a-zA-Z0-9][^720p]+\s*)`
	// if !yearFound {
	// 	// 	regexStr = regexStr + `{1,4}`
	// 	// } else {
	// 	// regexStr = regexStr + `{1,}`
	// }

	regexStr = regexStr + `{1,}`

	re = regexp.MustCompile(regexStr)
	result = re.FindString(result)
	// result = strings.ToLower(result)
	return strings.Trim(result, " ")
}

func getNameFromFolderName(folderName string) (string, bool) {
	if idx := findYearIndexFromString(folderName, true); idx == -1 {
		return "", false
	}

	return parseName(folderName), true
}

// func getYear(str string) string {
// 	re := regexp.MustCompile(`(20\d{2})|(19\d{2})`)
// 	result := re.FindString(str)
// 	return result
// }

// -1 - not found
func findYearIndexFromString(name string, strict bool) int {
	actualYear, _, _ := time.Now().Date()
	filter := regexYearString(actualYear)

	regexStr := fmt.Sprintf(`(\(%s\))|(\(19\d{2}\))`, filter)

	if idx := regexFindIndex(regexStr, name); idx != -1 {
		return idx + 1
	}
	if strict {
		return -1
	}
	regexStr = fmt.Sprintf(`(%s)|(19\d{2})`, filter)
	idx := regexFindIndex(regexStr, name)
	return idx
}

func regexFindIndex(regStr, str string) int {
	re := regexp.MustCompile(regStr)
	index := re.FindAllStringIndex(str, -1)
	if len(index) == 0 {
		return -1
	}
	return index[len(index)-1][0]
}

func regexYearString(year int) string {
	actualYear := year - 2000

	tens := actualYear / 10
	tensStr := "[0]"

	if tens > 2 {
		tensStr = "[0-" + strconv.Itoa(tens-1) + "]"
	}

	prf := strconv.Itoa(tens)

	lastYear := strconv.Itoa(actualYear % 10)

	return fmt.Sprintf(`(20(%s\d{1}|%s[0-%s]))`, tensStr, prf, lastYear)
}
