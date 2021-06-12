package app

import (
	"fmt"
	"github.com/adminoid/vuego/internal/config"
)

func Run(config *config.Config) {
	fmt.Println(config.DatabaseURL)
}
