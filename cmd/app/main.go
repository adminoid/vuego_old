package main

import (
	"github.com/adminoid/vuego/internal/app"
)

const configPath = "configs/main"

func main() {
	app.Run(configPath)
}
