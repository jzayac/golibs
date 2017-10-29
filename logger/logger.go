package logger

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Debug *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger

	green = "\x1b[32;1m"
	yelow = "\x1b[33;1m"
	blue  = "\x1b[34;1m"
	red   = "\x1b[31;1m"
	cyan  = "\x1b[36;1m"
	def   = "\x1b[0m"

	file = ""
)

const (
	standardLogFlags = log.Ldate | log.Ltime | log.Lmicroseconds
)

func init() {
	debugHandle := ioutil.Discard
	infoHandle := os.Stdout
	warningHandle := os.Stdout
	errorHandle := os.Stderr
	if os.Getenv("GO_ENV") == "develop" {
		// https://github.com/gopher-net/docker-ovs-plugin/blob/master/ovs/logging.go
		debugHandle = os.Stdout
	} else {
		green = ""
		yelow = ""
		blue = ""
		red = ""
		cyan = ""
		def = ""
	}

	// https://groups.google.com/forum/#!topic/Golang-nuts/nluStAtr8NA
	Debug = log.New(debugHandle, blue+"DEBUG:"+def+" ", standardLogFlags)
	Info = log.New(infoHandle, green+"INFO:"+def+"  ", standardLogFlags)
	Warn = log.New(warningHandle, yelow+"WARN:"+def+"  ", standardLogFlags)
	Error = log.New(errorHandle, red+"ERROR:"+def+" ", standardLogFlags)
}

func SetFile(f string) {
	file = f
}

func DebugLog(message string, method string) {
	Debug.Printf("%s%s:%s%s| %s", cyan, file, method, def, message)
	// Debug.Printf("%s%s:%s%s| %+v\n", cyan, file, method, def, message)
}

func InfoLog(message string, method string) {
	Info.Printf("%s%s:%s%s| %s\n", cyan, file, method, def, message)
}

func WarnLog(message string, method string) {
	Warn.Printf("%s%s:%s%s| %s\n", cyan, file, method, def, message)
}

func ErrorLog(message string, method string) {
	Error.Printf("%s%s:%s%s| %s\n", cyan, file, method, def, message)
}

func D(message string, method string) {
	DebugLog(message, method)
}

func I(message string, method string) {
	InfoLog(message, method)
}

func W(message string, method string) {
	WarnLog(message, method)
}

func E(message string, method string) {
	WarnLog(message, method)
}
