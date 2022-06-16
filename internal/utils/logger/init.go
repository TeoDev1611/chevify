package logger

import (
	"fmt"

	"github.com/TeoDev1611/batman/log"
)

func init() {
	log.Config.AppName = "chevify"
	log.Config.FileToLog = "chevifylogs.log"
	err := log.Init()
	if err != nil {
		panic(err)
	}
}

func Info(msg string) {
	string := fmt.Sprintf("CHEVIFY -> %s", msg)
	log.Info(string)
}

func Warn(msg string) {
	string := fmt.Sprintf("CHEVIFY -> %s", msg)
	log.Warning(string)
}

func Error(msg string) {
	string := fmt.Sprintf("CHEVIFY -> %s", msg)
	log.Error(string)
}

func Fatal(msg string) {
	string := fmt.Sprintf("CHEVIFY -> %s", msg)
	log.Fatal(string)
}
