package app

import (
	"fmt"
)

func Run(config *Config) {
	fmt.Println(config.DatabaseURL)
}