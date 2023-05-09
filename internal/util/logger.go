package util

import (
	"io"
	"log"
	"os"
)

func InitLog() {
	// 屏幕输出
	stdout := os.Stdout
	// 文件输出
	logfile, err := os.OpenFile("log.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create log file failed: %v", err)
	}

	logwriter := io.MultiWriter(stdout, logfile)
	log.SetOutput(logwriter)
}
