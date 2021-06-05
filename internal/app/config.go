package app

import (
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"log"
)

// Config ...
type Config struct {
	LogLevel    string
	DatabaseURL string
}

// NewConfig ...
func NewConfig() *Config {
	env := getEnv()
	config := &Config{}
	err := mapstructure.Decode(env, &config)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%+v\n", config)
	return config
}

func getEnv() map[string]string {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}
	return myEnv
}
