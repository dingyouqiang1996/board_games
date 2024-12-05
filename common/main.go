package main

import (
	"common/config"
	"common/logs"
)

func main() {
	config.InitConfig("./application.yml")
	logs.InitLog("test")
	logs.Debug("1111")
	logs.Info("2222")
	logs.Error("3333")
	logs.Fatal("4444")
}