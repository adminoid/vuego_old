package config

import (
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"log"
	"os"
)

var BasePath string = "$GOPATH/src/vuego/.env"

// Config ...
type Config struct {
	LogLevel    string
	DatabaseURL string
	BindAddr	string
}

// NewConfig ...
func NewConfig() Config {
	env := getEnv()
	config := &Config{}
	err := mapstructure.Decode(env, &config)
	if err != nil {
		log.Fatal(err)
	}
	return *config
}

func getEnv() map[string]string {
	var myEnv map[string]string
	myEnv, err := godotenv.Read(os.ExpandEnv(BasePath))
	if err != nil {
		log.Fatal(err)
	}
	return myEnv
}
