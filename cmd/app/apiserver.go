package main

import (
	"github.com/adminoid/vuego/internal/apiserver"
	"github.com/adminoid/vuego/internal/config"
	"log"
)

func main() {
	s := apiserver.New(config.NewConfig(""))
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
