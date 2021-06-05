package main

import (
	"github.com/adminoid/vuego/internal/app"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}
	dbUrl := myEnv["DATABASE"]
	app.Run(dbUrl)
}
