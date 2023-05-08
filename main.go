package main

import (
	"fmt"
	"log"
	"net/http"
	"openai/bootstrap"
	"openai/internal/config"
	"openai/internal/handler"
	"os"
)

func init() {

}

func main() {
	r := bootstrap.New()

	// 微信消息处理
	r.POST("/", handler.ReceiveMsg)
	// 用于公众号自动验证
	r.GET("/", handler.WechatCheck)
	// 用于测试 curl "http://127.0.0.1:$PORT/test"
	r.GET("/test", handler.Test)
	// r.GET("/", handler.Test)

	// 设置日志
	if !config.Debug {
		SetLog()
	}

	fmt.Printf("start process success at %s:%s \n", config.Http.Addr, config.Http.Port)
	if err := http.ListenAndServe(config.Http.Addr+":"+config.Http.Port, r); err != nil {
		panic(err)
	}
}

func SetLog() {
	file := "log.log"
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	fmt.Println("查看日志请使用 tail -f " + file)
}
