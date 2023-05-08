package util

import (
	"io"
	"log"
	"os"
)

var (
	Logger *log.Logger
)

func init() {
	// 屏幕输出
	stdout := os.Stdout
	// 文件输出
	logfile, err := os.OpenFile("log.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create log file failed: %v", err)
	}

	Logger = log.New(io.MultiWriter(stdout, logfile), "", log.Lshortfile|log.LstdFlags)
}
