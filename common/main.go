package main

import (
	"common/config"
	"common/logs"

	"github.com/charmbracelet/log"
)

func main() {
	config.InitConfig("./application.yml")
	logs.InitLog("test")
	log.Info("Hello, World!")
	log.Info("Hello, World!")
	log.Info("Hello, World!")
}