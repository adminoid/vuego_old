package main

import (
	"github.com/adminoid/vuego/internal/api"
	"github.com/adminoid/vuego/internal/config"
)

func main() {
	api.Run(config.NewConfig())
}
