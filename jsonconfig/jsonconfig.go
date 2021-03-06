package jsonconfig

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Parser must implement ParseJSON
type Parser interface {
	ParseJSON([]byte) error
	SetEnv(string)
}

// Load the JSON config file
// func Load(configFile string, p Parser) {
func Load(p Parser) {
	var err error
	var input = io.ReadCloser(os.Stdin)
	enviroment := os.Getenv("GOLANG_ENV")
	if enviroment != "DEVELOP" {
		enviroment = "PROD"
	}

	p.SetEnv(enviroment)

	enviroment = strings.ToLower(enviroment)
	separator := string(os.PathSeparator)
	configFile := "_config" + separator + enviroment + separator + "config.json"
	if input, err = os.Open(configFile); err != nil {
		log.Fatalln(err)
	}

	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the config
	if err := p.ParseJSON(jsonBytes); err != nil {
		log.Fatalln("Could not parse %q: %v", configFile, err)
	}
}
