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
	log.Debug("Hello, World!") // 这里我的调用没有按照课程的来，所以等级为info，打印不出debug的日志
	// log.Fatal("Hello, World!")
	log.Error("Hello, World!")
	log.Info("Hello, World!")
	a := 1
	go func() {
		log.Info(a)
	}()
	select {}
}
