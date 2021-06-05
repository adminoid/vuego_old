package app

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
)

// Config ...
type Config struct {
	LogLevel    string
	DatabaseURL string
}

// NewConfig ...
func NewConfig() *Config {

	env := getEnv()

	fmt.Println(reflect.TypeOf(env))

	jsonString, err := json.MarshalIndent(env, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(jsonString))
	fmt.Println(env["DATABASE"])

	os.Exit(1)

	return &Config{
		LogLevel: "debug",
	}
}

func getEnv() map[string]string {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}
	return myEnv
}
