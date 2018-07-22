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

	namespace = ""
)

type Logger struct {
	namespace string
}

func (l *Logger) Debug(message, method string) {
	Debug.Printf("%s%s:%s%s| %s", cyan, l.namespace, method, def, message)
}

func (l *Logger) Info(message, method string) {
	Info.Printf("%s%s:%s%s| %s\n", cyan, l.namespace, method, def, message)
}

func (l *Logger) Warn(message, method string) {
	Warn.Printf("%s%s:%s%s| %s\n", cyan, l.namespace, method, def, message)
}

func (l *Logger) Error(message, method string) {
	Error.Printf("%s%s:%s%s| %s\n", cyan, l.namespace, method, def, message)
}

const (
	standardLogFlags = log.Ldate | log.Ltime | log.Lmicroseconds
)

func init() {
	debugHandle := ioutil.Discard
	infoHandle := os.Stdout
	warningHandle := os.Stdout
	errorHandle := os.Stderr
	if os.Getenv("GOLANG_ENV") == "DEVELOP" || os.Getenv("GOLANG_ENV") == "TEST" {
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

func SetApp(f string) {
	namespace = f
}

func SetLogger(namespaceInput string) Logger {
	return Logger{namespace: namespaceInput}
}

func DebugLog(message string) {
	Debug.Printf("%s%s%s| %s", cyan, namespace, def, message)
	// Debug.Printf("%s%s:%s%s| %+v\n", cyan, namespace, message)
}

func InfoLog(message string) {
	Info.Printf("%s%s%s| %s\n", cyan, namespace, def, message)
}

func WarnLog(message string) {
	Warn.Printf("%s%s%s| %s\n", cyan, namespace, def, message)
}

func ErrorLog(message string) {
	Error.Printf("%s%s%s| %s\n", cyan, namespace, def, message)
}
