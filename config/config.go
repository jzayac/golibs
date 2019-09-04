package config

import (
	"encoding/json"
	"fmt"

	"github.com/jzayac/golibs/database"
	"github.com/jzayac/golibs/jsonconfig"
)

func init() {
	jsonconfig.Load(config)
}

type JsonFileConfig struct {
	Port     string
	Database database.Info `json:"pgDatabase"`
	Services Services
}

type Services struct {
	Socket    Service
	Tmdb      Service
	ParseName Service
}

type Service struct {
	Name     string
	Protocol string
	Port     int
}

func (sr Service) GetUrl() string {
	return fmt.Sprintf("%s://%s:%d", sr.Protocol, sr.Name, sr.Port)
}

var config = &configuration{}

type configuration struct {
	App JsonFileConfig
	Env string
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

func (c *configuration) SetEnv(env string) {
	c.Env = env
}

func GetDatabaseInfo() database.Info {
	return config.App.Database
}

func GetConfig() JsonFileConfig {
	return config.App
}

func IsDevelop() bool {
	return config.Env == "DEVELOP"
}

func GetEnv() string {
	return config.Env
}
