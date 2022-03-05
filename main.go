package main

import (
	"github.com/adiatma85/go-tutorial-gorm/src"
	"github.com/adiatma85/go-tutorial-gorm/src/config"
)

func main() {
	config.Initialize()
	src.Run()
}
