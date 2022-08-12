package config

import (
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"log"
	"os"
)

var BasePath string = "$GOPATH/src/vuego/"

// Config ...
type Config struct {
	DbUser string
	DbPwd string
	DbName string
	DbHost string
	LogLevel    string
	BindAddr	string
}

// NewConfig ...
func NewConfig(envFile string) Config {
	env := getEnv(envFile)
	config := &Config{}
	err := mapstructure.Decode(env, &config)
	if err != nil {
		log.Fatal(err)
	}
	return *config
}

func getEnv(envFile string) map[string]string {
	var myEnv map[string]string
	envPath := ".env"
	if len(envFile) > 0 {
		envPath = BasePath + envFile
	}
	myEnv, err := godotenv.Read(os.ExpandEnv(BasePath + envPath))
	if err != nil {
		log.Fatal(err)
	}
	return myEnv
}
